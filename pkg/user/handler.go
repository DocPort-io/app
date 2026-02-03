package user

import (
	"app/pkg/platform/handler"
	"app/pkg/platform/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	authMiddleware *middleware.AuthMiddleware
}

func NewHandler(authMiddleware *middleware.AuthMiddleware) *Handler {
	return &Handler{
		authMiddleware: authMiddleware,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(h.authMiddleware.Authenticate)
		r.Get("/me", h.GetMe)
	})
}

// GetMe godoc
//
//	@summary	Get current user information
//	@tags		users
//	@accept		json
//	@produce	json
//	@param		Authorization	header	string	true	"Authorization header in the format 'Bearer <token>'"
//	@success	200	{object}	UserResponse
//	@router		/api/v1/users/me [get]
func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	preferredUsername := r.Context().Value("AUTH_PREFERRED_USERNAME").(string)

	user := User{
		ID: preferredUsername,
	}

	handler.WriteJson(w, http.StatusOK, user.ToResponse())
}
