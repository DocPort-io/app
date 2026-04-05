package auth

import (
	"app/pkg/platform/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type Authenticator struct {
	config config.AuthConfig
	cache  *jwk.Cache
}

func NewAuthenticator(config config.AuthConfig) (*Authenticator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	httpClient := httprc.NewClient()

	cache, err := jwk.NewCache(ctx, httpClient)
	if err != nil {
		log.Printf("error creating jwk cache: %v", err)
		return nil, err
	}

	if err := cache.Register(ctx, config.JWKSUrl); err != nil {
		log.Printf("error registering JWKS: %v", err)
		return nil, err
	}

	return &Authenticator{config: config, cache: cache}, nil
}

func (m *Authenticator) Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName == "OpenIdConnect" {
		err, unverifiedToken := checkToken(input.RequestValidationInput.Request)
		if err != nil {
			return err
		}

		keySet, err := m.cache.Lookup(ctx, m.config.JWKSUrl)
		if err != nil {
			log.Printf("error looking up jwk set: %v", err)
			return errors.New("internal server error")
		}

		token, err := jwt.ParseString(unverifiedToken, jwt.WithKeySet(keySet))
		if errors.Is(err, jwt.TokenExpiredError()) {
			return errors.New("token expired")
		}
		if err != nil {
			log.Printf("error parsing token: %v", err)
			return errors.New("invalid token")
		}

		var scopeString string
		if err := token.Get("scope", &scopeString); err != nil {
			return errors.New("invalid scope")
		}

		scopes := strings.Split(scopeString, " ")

		if len(m.config.Scopes) > 0 {
			for _, scope := range m.config.Scopes {
				if !slices.Contains(scopes, scope) {
					return errors.New(fmt.Sprintf("missing scope %s", scope))
				}
			}
		}
	}

	return nil
}

// GetUnverifiedToken parses a JWT from the request and extracts claims without
// verifying the token's signature. This function MUST ONLY be called on
// requests that have already been authenticated and whose JWT has already
// been validated (e.g. by Authenticator.Authenticate). Do not use this
// helper to perform authentication or authorization decisions on
// unauthenticated requests.
func GetUnverifiedToken(request *http.Request) (error, TokenContext) {
	err, unverifiedToken := checkToken(request)
	if err != nil {
		return err, TokenContext{}
	}

	token, err := jwt.ParseString(unverifiedToken, jwt.WithVerify(false))
	if errors.Is(err, jwt.TokenExpiredError()) {
		return errors.New("token expired"), TokenContext{}
	}
	if err != nil {
		log.Printf("error parsing token: %v", err)
		return errors.New("invalid token"), TokenContext{}
	}

	var scopeString string
	if err := token.Get("scope", &scopeString); err != nil {
		return errors.New("invalid scope"), TokenContext{}
	}

	scopes := strings.Split(scopeString, " ")

	subject, ok := token.Subject()
	if !ok {
		return errors.New("invalid token"), TokenContext{}
	}

	issuer, ok := token.Issuer()
	if !ok {
		return errors.New("invalid token"), TokenContext{}
	}

	var name string
	err = token.Get("name", &name)
	if err != nil {
		return errors.New("invalid token"), TokenContext{}
	}

	var givenName string
	err = token.Get("given_name", &givenName)
	if err != nil {
		return errors.New("invalid token"), TokenContext{}
	}

	var familyName string
	err = token.Get("family_name", &familyName)
	if err != nil {
		return errors.New("invalid token"), TokenContext{}
	}

	var email string
	err = token.Get("email", &email)
	if err != nil {
		return errors.New("invalid token"), TokenContext{}
	}

	var emailVerified bool
	err = token.Get("email_verified", &emailVerified)
	if err != nil {
		return errors.New("invalid token"), TokenContext{}
	}

	var preferredUsername string
	err = token.Get("preferred_username", &preferredUsername)
	if err != nil {
		return errors.New("invalid token"), TokenContext{}
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

	return nil, tokenContext
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

func checkToken(r *http.Request) (error, string) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return errors.New("missing authorization header"), ""
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		return errors.New("invalid authorization header"), ""
	}
	return nil, strings.TrimPrefix(auth, "Bearer ")
}
