package middleware

import (
	"app/pkg/platform/config"
	"app/pkg/platform/handler"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

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

		_, err = jwt.ParseString(unverifiedToken, jwt.WithKeySet(keySet))
		if errors.Is(err, jwt.TokenExpiredError()) {
			handler.WriteError(w, http.StatusUnauthorized, "token expired")
			return
		}
		if err != nil {
			log.Printf("error parsing token: %v", err)
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
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
