package user

import (
	"app/pkg/api"
	"app/pkg/platform/auth"
	"app/pkg/platform/handler"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	openapitypes "github.com/oapi-codegen/runtime/types"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/v1/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)

		r.Route("/me", func(r chi.Router) {
			r.Get("/", h.GetMe)
			r.Get("/token-info", h.GetMyTokenInfo)
		})

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", h.GetUser)
		})
	})
}

func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	err, tokenContext := auth.GetUnverifiedToken(r)
	if err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	user, err := h.service.GetByKeycloakReference(r.Context(), tokenContext.Subject)
	if errors.Is(err, ErrUserNotFound) {
		writeUserNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toUserResponse(user))
}

func (h *Handler) GetMyTokenInfo(w http.ResponseWriter, r *http.Request) {
	err, tokenContext := auth.GetUnverifiedToken(r)
	if err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toTokenInfoResponse(tokenContext))
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := parseUserId(r)
	if err != nil {
		writeInvalidUserIdError(w)
		return
	}

	user, err := h.service.GetById(r.Context(), userId)
	if errors.Is(err, ErrUserNotFound) {
		writeUserNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toUserResponse(user))
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req api.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	user, err := h.service.CreateUser(r.Context(), CreateUserRequest{
		Name:          req.Name,
		Email:         string(req.Email),
		EmailVerified: req.EmailVerified,
	})
	if errors.Is(err, ErrUserAlreadyExists) {
		writeUserAlreadyExistsError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toUserResponse(user))
}

func writeInvalidUserIdError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusBadRequest, "invalid user id")
}

func writeUserNotFoundError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusNotFound, "user not found")
}

func writeUserAlreadyExistsError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusConflict, "user already exists")
}

func parseUserId(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
}

func toUserResponse(u User) api.UserResponse {
	return api.UserResponse{
		Id:            u.ID,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
		Name:          u.Name,
		Email:         openapitypes.Email(u.Email),
		EmailVerified: u.EmailVerified,
	}
}

func toTokenInfoResponse(tokenContext auth.TokenContext) api.TokenInfoResponse {
	return api.TokenInfoResponse{
		Subject: tokenContext.Subject,
	}
}
