package user

import (
	"app/pkg/platform/handler"
	"app/pkg/platform/middleware"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service        Service
	validate       *validator.Validate
	authMiddleware *middleware.AuthMiddleware
}

func NewHandler(service Service, authMiddleware *middleware.AuthMiddleware) *Handler {
	return &Handler{
		service:        service,
		validate:       validator.New(),
		authMiddleware: authMiddleware,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(h.authMiddleware.Authenticate)
		r.Post("/", h.CreateUser)

		r.Route("/me", func(r chi.Router) {
			r.Get("/", h.GetMe)
			r.Get("/external-auths", h.ListMyExternalAuths)
			r.Get("/token-info", h.GetMyTokenInfo)
		})

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", h.GetUser)
			r.Get("/external-auths", h.ListUserExternalAuths)
			r.Post("/external-auths", h.CreateUserExternalAuth)
		})
	})
}

// GetMe godoc
//
//	@summary	Get current user
//	@tags		users
//	@accept		json
//	@produce	json
//	@success	200	{object}	UserResponse
//	@failure	404	{object}	handler.ErrorResponse
//	@failure	500	{object}	handler.ErrorResponse
//	@security	OAuth2AccessCode
//	@router		/api/v1/users/me [get]
func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	tokenContext := middleware.GetTokenContextFromContext(r.Context())

	user, err := h.service.GetByProvider(r.Context(), tokenContext.Issuer, tokenContext.Subject)
	if errors.Is(err, ErrUserNotFound) {
		writeUserNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, user.ToResponse())
}

func (h *Handler) GetMyTokenInfo(w http.ResponseWriter, r *http.Request) {
	tokenContext := middleware.GetTokenContextFromContext(r.Context())
	handler.WriteJson(w, http.StatusOK, toTokenInfoResponse(tokenContext))
}

// ListMyExternalAuths godoc
//
//	@summary	Get external auths for current user
//	@tags		users
//	@accept		json
//	@produce	json
//	@success	200	{object}	ListExternalAuthResponse
//	@failure	404	{object}	handler.ErrorResponse
//	@failure	500	{object}	handler.ErrorResponse
//	@security	OAuth2AccessCode
//	@router		/api/v1/users/me/external-auths [get]
func (h *Handler) ListMyExternalAuths(w http.ResponseWriter, r *http.Request) {
	tokenContext := middleware.GetTokenContextFromContext(r.Context())

	user, err := h.service.GetByProvider(r.Context(), tokenContext.Issuer, tokenContext.Subject)
	if errors.Is(err, ErrUserNotFound) {
		writeUserNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	externalAuths, err := h.service.ListExternalAuths(r.Context(), user.ID)
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, ToListExternalAuthResponse(externalAuths))
}

// GetUser godoc
//
//	@summary	Get a user
//	@tags		users
//	@accept		json
//	@produce	json
//	@param		userId	path		uint	true	"User identifier"
//	@success	200		{object}	UserResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	404		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@security	OAuth2AccessCode
//	@router		/api/v1/users/{userId} [get]
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

	handler.WriteJson(w, http.StatusOK, user.ToResponse())
}

// ListUserExternalAuths godoc
//
//	@summary	Get external auths for a user
//	@tags		users
//	@accept		json
//	@produce	json
//	@param		userId	path		uint	true	"User identifier"
//	@success	200		{object}	ListExternalAuthResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	404		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@security	OAuth2AccessCode
//	@router		/api/v1/users/{userId}/external-auths [get]
func (h *Handler) ListUserExternalAuths(w http.ResponseWriter, r *http.Request) {
	userId, err := parseUserId(r)
	if err != nil {
		writeInvalidUserIdError(w)
		return
	}

	externalAuths, err := h.service.ListExternalAuths(r.Context(), userId)
	if errors.Is(err, ErrUserNotFound) {
		writeUserNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, ToListExternalAuthResponse(externalAuths))
}

// CreateUser godoc
//
//	@summary	Create a user
//	@tags		users
//	@accept		json
//	@produce	json
//	@param		request	body		CreateUserRequest	true	"Create a user"
//	@success	201		{object}	UserResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	409		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@router		/api/v1/users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	user, err := h.service.CreateUser(r.Context(), req)
	if errors.Is(err, ErrUserAlreadyExists) {
		writeUserAlreadyExistsError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, user.ToResponse())
}

// CreateUserExternalAuth godoc
//
//	@summary	Create a user external auth
//	@tags		users
//	@accept		json
//	@produce	json
//	@param		request	body		CreateUserRequest	true	"Create a user"
//	@success	201		{object}	UserResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	409		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@router		/api/v1/users/{userId}/external-auths [post]
func (h *Handler) CreateUserExternalAuth(w http.ResponseWriter, r *http.Request) {
	id, err := parseUserId(r)
	if err != nil {
		writeInvalidUserIdError(w)
		return
	}

	var req CreateExternalAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	externalAuth, err := h.service.CreateExternalAuth(r.Context(), id, req)
	if errors.Is(err, ErrExternalAuthAlreadyExists) {
		writeExternalAuthAlreadyExistsError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusCreated, externalAuth.ToResponse())
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

func writeExternalAuthAlreadyExistsError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusConflict, "external auth already exists")
}

func parseUserId(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
}

type TokenInfoResponse struct {
	Subject string `json:"sub"`
}

func toTokenInfoResponse(tokenContext middleware.TokenContext) TokenInfoResponse {
	return TokenInfoResponse{
		Subject: tokenContext.Subject,
	}
}
