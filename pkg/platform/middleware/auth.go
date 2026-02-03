package middleware

import (
	"app/pkg/platform/handler"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/zitadel/oidc/v3/pkg/client/rs"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

const JWKS_URI = `https://keycloak.docport.io/realms/docport-dev`

type AuthMiddleware struct {
	provider rs.ResourceServer
}

func NewAuthMiddleware() (*AuthMiddleware, error) {
	ctx := context.Background()

	provider, err := rs.NewResourceServerClientCredentials(ctx, JWKS_URI, "docport-dev", "cH0rHPr1w4hFaQCj44Rtxtv5uOWHfmA3")
	if err != nil {
		return nil, err
	}

	return &AuthMiddleware{provider: provider}, nil
}

func (am *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok, token := checkToken(w, r)
		if !ok {
			return
		}

		log.Printf("token: %s", token)

		resp, err := rs.Introspect[*oidc.IntrospectionResponse](r.Context(), am.provider, token)
		if err != nil {
			log.Printf("error in introspection: %v", err)
			handler.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		log.Printf(resp.PreferredUsername)

		//verifiedToken, err := jwt.ParseRequest(r, jwt.WithHeaderKey("Authorization"), jwt.WithKeySet(jwkSet))
		//if err != nil {
		//	log.Printf("failed to parse JWT token: %v", err)
		//	handler.WriteError(w, http.StatusUnauthorized, "invalid token")
		//	return
		//}
		//
		//var preferredUsername string
		//err = verifiedToken.Get("preferred_username", &preferredUsername)
		//if err != nil {
		//	log.Printf("failed to get preferred username: %v", err)
		//	handler.WriteError(w, http.StatusUnauthorized, "missing preferred_username in token")
		//	return
		//}

		ctx := context.WithValue(r.Context(), "AUTH_PREFERRED_USERNAME", resp.PreferredUsername)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func checkToken(w http.ResponseWriter, r *http.Request) (bool, string) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		handler.WriteError(w, http.StatusUnauthorized, "missing authorization header")
		return false, ""
	}
	if !strings.HasPrefix(auth, oidc.PrefixBearer) {
		handler.WriteError(w, http.StatusUnauthorized, "invalid authorization header")
		return false, ""
	}
	return true, strings.TrimPrefix(auth, oidc.PrefixBearer)
}

func fetchJwkSet() (jwk.Set, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jwkSet, err := jwk.Fetch(ctx, JWKS_URI)
	if err != nil {
		return nil, fmt.Errorf("failed to load JWK set: %w", err)
	}
	return jwkSet, nil
}
