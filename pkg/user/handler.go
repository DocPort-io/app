package user

import (
	"app/pkg/platform/handler"
	"app/pkg/platform/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Get("/me", h.GetMe)
	})
}

// GetMe godoc
//
//	@summary	Get current user information
//	@tags		users
//	@accept		json
//	@produce	json
//	@success	200	{object}	UserResponse
//	@router		/api/v1/users/me [get]
func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID: "",
	}

	handler.WriteJson(w, http.StatusOK, user.ToResponse())
}
