package middleware

import (
	"app/pkg/platform/config"
	"app/pkg/platform/handler"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type contextKey string

const tokenContextKey contextKey = "token_context"

type AuthMiddleware struct {
	config config.Config
	cache  *jwk.Cache
}

func NewAuthMiddleware(config config.Config) (*AuthMiddleware, error) {
	ctx := context.Background()

	httpClient := httprc.NewClient()

	cache, err := jwk.NewCache(ctx, httpClient)
	if err != nil {
		log.Printf("error creating jwk cache: %v", err)
		return nil, err
	}

	if err := cache.Register(ctx, config.Auth.JWKSUrl); err != nil {
		log.Printf("error registering JWKS: %v", err)
		return nil, err
	}

	return &AuthMiddleware{config: config, cache: cache}, nil
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok, unverifiedToken := checkToken(w, r)
		if !ok {
			return
		}

		keySet, err := m.cache.Lookup(r.Context(), m.config.Auth.JWKSUrl)
		if err != nil {
			log.Printf("error looking up jwk set: %v", err)
			handler.WriteError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		token, err := jwt.ParseString(unverifiedToken, jwt.WithKeySet(keySet))
		if errors.Is(err, jwt.TokenExpiredError()) {
			handler.WriteError(w, http.StatusUnauthorized, "token expired")
			return
		}
		if err != nil {
			log.Printf("error parsing token: %v", err)
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		var scopeString string
		if err := token.Get("scope", &scopeString); err != nil {
			handler.WriteError(w, http.StatusUnauthorized, "invalid scope")
			return
		}

		scopes := strings.Split(scopeString, " ")

		if len(m.config.Auth.Scopes) > 0 {
			for _, scope := range m.config.Auth.Scopes {
				if !slices.Contains(scopes, scope) {
					handler.WriteError(w, http.StatusUnauthorized, fmt.Sprintf("missing scope %s", scope))
					return
				}
			}
		}

		subject, ok := token.Subject()
		if !ok {
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		issuer, ok := token.Issuer()
		if !ok {
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		var name string
		err = token.Get("name", &name)
		if err != nil {
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		var givenName string
		err = token.Get("given_name", &givenName)
		if err != nil {
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		var familyName string
		err = token.Get("family_name", &familyName)
		if err != nil {
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		var email string
		err = token.Get("email", &email)
		if err != nil {
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		var emailVerified bool
		err = token.Get("email_verified", &emailVerified)
		if err != nil {
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		var preferredUsername string
		err = token.Get("preferred_username", &preferredUsername)
		if err != nil {
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		tokenContext := TokenContext{
			Subject:           subject,
			Issuer:            issuer,
			Name:              name,
			GivenName:         givenName,
			FamilyName:        familyName,
			Email:             email,
			EmailVerified:     emailVerified,
			PreferredUsername: preferredUsername,
			Scopes:            scopes,
		}

		ctx := context.WithValue(r.Context(), tokenContextKey, tokenContext)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func GetTokenContextFromContext(ctx context.Context) TokenContext {
	tokenContext := ctx.Value(tokenContextKey)
	if tokenContext == nil {
		log.Fatalf("token context not found in context")
	}
	return tokenContext.(TokenContext)
}

type TokenContext struct {
	Subject           string
	Issuer            string
	Name              string
	GivenName         string
	FamilyName        string
	Email             string
	EmailVerified     bool
	PreferredUsername string
	Scopes            []string
}

func checkToken(w http.ResponseWriter, r *http.Request) (bool, string) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		handler.WriteError(w, http.StatusUnauthorized, "missing authorization header")
		return false, ""
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		handler.WriteError(w, http.StatusUnauthorized, "invalid authorization header")
		return false, ""
	}
	return true, strings.TrimPrefix(auth, "Bearer ")
}
