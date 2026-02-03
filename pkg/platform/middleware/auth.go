package middleware

import (
	"app/pkg/platform/config"
	"app/pkg/platform/handler"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/zitadel/oidc/v3/pkg/client/rs"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

type AuthMiddleware struct {
	provider rs.ResourceServer
}

func NewAuthMiddleware(cfg config.Config) (*AuthMiddleware, error) {
	ctx := context.Background()

	provider, err := rs.NewResourceServerClientCredentials(ctx, cfg.Auth.Issuer, cfg.Auth.ClientId, cfg.Auth.ClientSecret)
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

		if resp.Active == false {
			handler.WriteError(w, http.StatusUnauthorized, "token is not active")
			return
		}

		data, err := json.Marshal(resp)
		if err != nil {
			log.Printf("error marshalling response: %v", err)
			handler.WriteError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		ctx := context.WithValue(r.Context(), "BEARER_AUTH_JSON", data)

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
