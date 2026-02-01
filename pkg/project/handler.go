package project

import (
	"app/pkg/platform/handler"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  Service
	validate *validator.Validate
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service, validate: validator.New()}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/projects", func(r chi.Router) {
		r.Get("/", h.List)
		r.Post("/", h.Create)

		r.Route("/{projectId}", func(r chi.Router) {
			r.Get("/", h.GetById)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
		})
	})
}

// GetById godoc
//
//	@summary	Get a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		projectId	path		uint	true	"Project identifier"
//	@success	200			{object}	ProjectResponse
//	@failure	400			{object}	handler.ErrorResponse
//	@failure	404			{object}	handler.ErrorResponse
//	@failure	500			{object}	handler.ErrorResponse
//	@router		/api/v1/projects/{projectId} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := parseProjectId(r)
	if err != nil {
		writeInvalidProjectIdError(w)
		return
	}

	project, err := h.service.GetById(r.Context(), id)
	if errors.Is(err, ErrProjectNotFound) {
		writeProjectNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, project.ToResponse())
}

// List godoc
//
//	@summary	Find all projects
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		limit	query		uint	false	"Max items per page (1-100)"
//	@param		offset	query		uint	false	"Items to skip before starting to collect the result set"
//	@success	200		{object}	ListProjectsResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@router		/api/v1/projects [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	limit, offset := handler.ParsePagination(r)

	projects, err := h.service.List(r.Context(), limit, offset)
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, ToListResponse(projects, limit, offset))
}

// Create godoc
//
//	@summary	Create a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		request	body		CreateProjectRequest	true	"Create a project"
//	@success	201		{object}	ProjectResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@router		/api/v1/projects [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	project, err := h.service.Create(r.Context(), req)
	if errors.Is(err, ErrProjectAlreadyExists) {
		writeProjectAlreadyExistsError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusCreated, project.ToResponse())
}

// Update godoc
//
//	@summary	Update a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		projectId	path		uint					true	"Project identifier"
//	@param		request		body		UpdateProjectRequest	true	"Update a project"
//	@success	200			{object}	ProjectResponse
//	@failure	400			{object}	handler.ErrorResponse
//	@failure	404			{object}	handler.ErrorResponse
//	@failure	500			{object}	handler.ErrorResponse
//	@router		/api/v1/projects/{projectId} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseProjectId(r)
	if err != nil {
		writeInvalidProjectIdError(w)
		return
	}

	var req UpdateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	project, err := h.service.Update(r.Context(), id, req)
	if errors.Is(err, ErrProjectNotFound) {
		writeProjectNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, project.ToResponse())
}

// Delete godoc
//
//	@summary	Delete a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		projectId	path	uint	true	"Project identifier"
//	@success	204
//	@failure	404	{object}	handler.ErrorResponse
//	@failure	500	{object}	handler.ErrorResponse
//	@router		/api/v1/projects/{projectId} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseProjectId(r)
	if err != nil {
		writeInvalidProjectIdError(w)
		return
	}

	err = h.service.Delete(r.Context(), id)
	if errors.Is(err, ErrProjectNotFound) {
		writeProjectNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeInvalidProjectIdError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusBadRequest, "invalid project id")
}

func writeProjectNotFoundError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusNotFound, "project not found")
}

func writeProjectAlreadyExistsError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusConflict, "project already exists")
}

func parseProjectId(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "projectId"), 10, 64)
}
