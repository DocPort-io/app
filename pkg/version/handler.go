package version

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
	r.Route("/versions", func(r chi.Router) {
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

// GetById godoc
//
//	@summary	Get a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		versionId	path		uint	true	"Version identifier"
//	@success	200			{object}	VersionResponse
//	@failure	400			{object}	handler.ErrorResponse
//	@failure	404			{object}	handler.ErrorResponse
//	@failure	500			{object}	handler.ErrorResponse
//	@router		/api/v1/versions/{versionId} [get]
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

	handler.WriteJson(w, http.StatusOK, version.ToResponse())
}

// List godoc
//
//	@summary	Find all versions
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		projectId	query		uint	false	"Project identifier"
//	@param		limit		query		uint	false	"Max items per page (1-100)"
//	@param		offset		query		uint	false	"Items to skip before starting to collect the result set"
//	@success	200			{object}	ListVersionsResponse
//	@failure	400			{object}	handler.ErrorResponse
//	@failure	500			{object}	handler.ErrorResponse
//	@router		/api/v1/versions [get]
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

	handler.WriteJson(w, http.StatusOK, ToListResponse(versions, limit, offset))
}

// Create godoc
//
//	@summary	Create a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		request	body		CreateVersionRequest	true	"Create a version"
//	@success	201		{object}	VersionResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@router		/api/v1/versions [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateVersionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	version, err := h.service.Create(r.Context(), req)
	if errors.Is(err, ErrVersionAlreadyExists) {
		writeVersionAlreadyExistsError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusCreated, version.ToResponse())
}

// Update godoc
//
//	@summary	Update a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		versionId	path		uint					true	"Version identifier"
//	@param		request		body		UpdateVersionRequest	true	"Update a version"
//	@success	200			{object}	VersionResponse
//	@failure	400			{object}	handler.ErrorResponse
//	@failure	404			{object}	handler.ErrorResponse
//	@failure	500			{object}	handler.ErrorResponse
//	@router		/api/v1/versions/{versionId} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	var req UpdateVersionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	version, err := h.service.Update(r.Context(), id, req)
	if errors.Is(err, ErrVersionNotFound) {
		writeVersionNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, version.ToResponse())
}

// Delete godoc
//
//	@summary	Delete a version
//	@tags		versions
//	@accept		json
//	@param		versionId	path	uint	true	"Version identifier"
//	@success	204
//	@failure	404	{object}	handler.ErrorResponse
//	@failure	500	{object}	handler.ErrorResponse
//	@router		/api/v1/versions/{versionId} [delete]
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

// AttachFile godoc
//
//	@summary	Attaches a file to a version
//	@tags		versions
//	@accept		json
//	@param		versionId	path	uint				true	"Version identifier"
//	@param		request		body	AttachFileRequest	true	"File to attach"
//	@success	204
//	@failure	400	{object}	handler.ErrorResponse
//	@failure	404	{object}	handler.ErrorResponse
//	@failure	500	{object}	handler.ErrorResponse
//	@router		/api/v1/versions/{versionId}/attach-file [patch]
func (h *Handler) AttachFile(w http.ResponseWriter, r *http.Request) {
	id, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	var req AttachFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	err = h.service.AttachFile(r.Context(), id, req)
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

// DetachFile godoc
//
//	@summary	Detach a file from a version
//	@tags		versions
//	@accept		json
//	@param		versionId	path	uint				true	"Version identifier"
//	@param		request		body	DetachFileRequest	true	"File to detach"
//	@success	204
//	@failure	400	{object}	handler.ErrorResponse
//	@failure	404	{object}	handler.ErrorResponse
//	@failure	500	{object}	handler.ErrorResponse
//	@router		/api/v1/versions/{versionId}/detach-file [patch]
func (h *Handler) DetachFile(w http.ResponseWriter, r *http.Request) {
	id, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	var req DetachFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	err = h.service.DetachFile(r.Context(), id, req)
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
