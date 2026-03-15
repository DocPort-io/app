package version

import (
	"app/pkg/api"
	"app/pkg/platform/handler"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/v1/versions", func(r chi.Router) {
		r.Get("/", h.List)
		r.Post("/", h.Create)

		r.Route("/{versionId}", func(r chi.Router) {
			r.Get("/", h.GetById)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
			r.Patch("/attach-file", h.AttachFile)
			r.Patch("/detach-file", h.DetachFile)
		})
	})
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	version, err := h.service.GetById(r.Context(), id)
	if errors.Is(err, ErrVersionNotFound) {
		writeVersionNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toVersionResponse(version))
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	limit, offset := handler.ParsePagination(r)

	projectId, err := parseProjectId(r)
	if err != nil {
		writeInvalidProjectIdError(w)
		return
	}

	versions, err := h.service.List(r.Context(), projectId, limit, offset)
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toListVersionsResponse(versions, limit, offset))
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req api.CreateVersionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	version, err := h.service.Create(r.Context(), CreateVersionRequest{
		Name:        req.Name,
		Description: req.Description,
		ProjectId:   req.ProjectId,
	})
	if errors.Is(err, ErrVersionAlreadyExists) {
		writeVersionAlreadyExistsError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusCreated, toVersionResponse(version))
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	var req api.UpdateVersionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	version, err := h.service.Update(r.Context(), id, UpdateVersionRequest{
		Name:        req.Name,
		Description: req.Description,
	})
	if errors.Is(err, ErrVersionNotFound) {
		writeVersionNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, toVersionResponse(version))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	err = h.service.Delete(r.Context(), id)
	if errors.Is(err, ErrVersionNotFound) {
		writeVersionNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) AttachFile(w http.ResponseWriter, r *http.Request) {
	id, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	var req api.AttachFileToVersionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	err = h.service.AttachFile(r.Context(), id, AttachFileRequest{
		FileID: req.FileId,
	})
	if errors.Is(err, ErrVersionNotFound) {
		writeVersionNotFoundError(w)
		return
	}
	if errors.Is(err, ErrVersionFileAlreadyAttached) {
		writeVersionFileAlreadyAttachedError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DetachFile(w http.ResponseWriter, r *http.Request) {
	id, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	var req api.DetachFileFromVersionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	err = h.service.DetachFile(r.Context(), id, DetachFileRequest{
		FileID: req.FileId,
	})
	if errors.Is(err, ErrVersionNotFound) {
		writeVersionNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeInvalidVersionIdError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusBadRequest, "invalid version id")
}

func writeVersionNotFoundError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusNotFound, "version not found")
}

func writeVersionAlreadyExistsError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusConflict, "version already exists")
}

func parseVersionId(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "versionId"), 10, 64)
}

func parseProjectId(r *http.Request) (*int64, error) {
	if !r.URL.Query().Has("projectId") {
		return nil, nil
	}

	projectId, err := strconv.ParseInt(r.URL.Query().Get("projectId"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &projectId, nil
}

func writeInvalidProjectIdError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusBadRequest, "invalid project id")
}

func writeVersionFileAlreadyAttachedError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusConflict, "version file already attached")
}

func toVersionResponse(v Version) api.VersionResponse {
	return api.VersionResponse{
		Id:          v.ID,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Name:        v.Name,
		Description: v.Description,
		ProjectId:   v.ProjectID,
	}
}

func toListVersionsResponse(versions []Version, limit, offset int64) api.ListVersionsResponse {
	items := make([]api.VersionResponse, len(versions))
	for i, version := range versions {
		items[i] = toVersionResponse(version)
	}
	return api.ListVersionsResponse{
		Limit:    limit,
		Offset:   offset,
		Versions: items,
	}
}
