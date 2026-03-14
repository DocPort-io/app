package project

import (
	"app/pkg/api"
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
	r.Route("/v1/projects", func(r chi.Router) {
		r.Get("/", h.List)
		r.Post("/", h.Create)

		r.Route("/{projectId}", func(r chi.Router) {
			r.Get("/", h.GetById)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
		})
	})
}

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

	handler.WriteJson(w, http.StatusOK, toProjectResponse(project))
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	limit, offset := handler.ParsePagination(r)

	projects, err := h.service.List(r.Context(), limit, offset)
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toListProjectsResponse(projects, limit, offset))
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req api.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	project, err := h.service.Create(r.Context(), CreateProjectRequest{
		Slug: req.Slug,
		Name: req.Name,
	})
	if errors.Is(err, ErrProjectAlreadyExists) {
		writeProjectAlreadyExistsError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusCreated, toProjectResponse(project))
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseProjectId(r)
	if err != nil {
		writeInvalidProjectIdError(w)
		return
	}

	var req api.UpdateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	project, err := h.service.Update(r.Context(), id, UpdateProjectRequest{
		Slug: req.Slug,
		Name: req.Name,
	})
	if errors.Is(err, ErrProjectNotFound) {
		writeProjectNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toProjectResponse(project))
}

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

func toProjectResponse(p Project) api.ProjectResponse {
	return api.ProjectResponse{
		Id:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Slug:      p.Slug,
		Name:      p.Name,
	}
}

func toListProjectsResponse(p []Project, limit, offset int64) api.ListProjectsResponse {
	items := make([]api.ProjectResponse, len(p))
	for i, project := range p {
		items[i] = toProjectResponse(project)
	}
	return api.ListProjectsResponse{
		Limit:    limit,
		Offset:   offset,
		Projects: items,
	}
}
