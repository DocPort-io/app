package file

import (
	"app/pkg/platform/handler"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/dustin/go-humanize"
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
	r.Route("/files", func(r chi.Router) {
		r.Get("/", h.List)
		r.Post("/", h.Create)

		r.Route("/{fileId}", func(r chi.Router) {
			r.Get("/", h.GetById)
			r.Post("/upload", h.Upload)
			r.Get("/download", h.Download)
			r.Delete("/", h.Delete)
		})
	})
}

// GetById godoc
//
//	@summary	Get a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		fileId	path		uint	true	"File identifier"
//	@success	200		{object}	FileResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	404		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@router		/api/v1/files/{fileId} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := parseFileId(r)
	if err != nil {
		writeInvalidFileIdError(w)
		return
	}

	file, err := h.service.GetById(r.Context(), id)
	if errors.Is(err, ErrFileNotFound) {
		writeFileNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, file.ToResponse())
}

// List godoc
//
//	@summary	Find all files
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		versionId	query		uint	false	"Version identifier"
//	@param		limit		query		uint	false	"Max items per page (1-100)"
//	@param		offset		query		uint	false	"Items to skip before starting to collect the result set"
//	@success	200			{object}	ListFilesResponse
//	@failure	400			{object}	handler.ErrorResponse
//	@failure	500			{object}	handler.ErrorResponse
//	@router		/api/v1/files [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	limit, offset := handler.ParsePagination(r)

	versionId, err := parseVersionId(r)
	if err != nil {
		writeInvalidVersionIdError(w)
		return
	}

	files, err := h.service.List(r.Context(), versionId, limit, offset)
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusOK, ToListResponse(files, limit, offset))
}

// Create godoc
//
//	@summary	Create a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		request	body		CreateFileRequest	true	"Create a file"
//	@success	201		{object}	FileResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@router		/api/v1/files [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		handler.WriteValidationError(w, err)
		return
	}

	file, err := h.service.Create(r.Context(), req)
	if errors.Is(err, ErrFileAlreadyExist) {
		writeFileAlreadyExistsError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusCreated, file.ToResponse())
}

// Upload godoc
//
//	@summary	Upload a file
//	@tags		files
//	@accept		multipart/form-data
//	@produce	json
//	@param		fileId	path		uint	true	"File identifier"
//	@param		file	formData	file	true	"File to upload"
//	@success	201		{object}	FileResponse
//	@failure	400		{object}	handler.ErrorResponse
//	@failure	404		{object}	handler.ErrorResponse
//	@failure	409		{object}	handler.ErrorResponse
//	@failure	500		{object}	handler.ErrorResponse
//	@router		/api/v1/files/{fileId}/upload [post]
func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	id, err := parseFileId(r)
	if err != nil {
		writeInvalidFileIdError(w)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1*humanize.GiByte)

	multipartFile, multipartFileHeader, err := r.FormFile("file")
	if err != nil {
		handler.WriteInvalidRequestPayloadError(w)
		return
	}

	req := UploadFileRequest{File: multipartFile, FileHeader: multipartFileHeader}

	uploadResult, err := h.service.UploadFile(r.Context(), id, req)
	if errors.Is(err, ErrFileNotFound) {
		writeFileNotFoundError(w)
		return
	}
	if errors.Is(err, ErrFileAlreadyComplete) {
		writeFileAlreadyCompleteError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	handler.WriteJson(w, http.StatusCreated, uploadResult.ToResponse())
}

// Download godoc
//
//	@summary	Download a file
//	@tags		files
//	@accept		json
//	@param		fileId	path	uint	true	"File identifier"
//	@success	200
//	@failure	400	{object}	handler.ErrorResponse
//	@failure	404	{object}	handler.ErrorResponse
//	@failure	500	{object}	handler.ErrorResponse
//	@router		/api/v1/files/{fileId}/download [get]
func (h *Handler) Download(w http.ResponseWriter, r *http.Request) {
	id, err := parseFileId(r)
	if err != nil {
		writeInvalidFileIdError(w)
		return
	}

	file, reader, err := h.service.Download(r.Context(), id)
	if errors.Is(err, ErrFileNotFound) {
		writeFileNotFoundError(w)
		return
	}
	if errors.Is(err, ErrFileNotComplete) {
		writeFileNotCompleteError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}
	defer func(reader io.ReadSeekCloser) {
		err := reader.Close()
		if err != nil {
			log.Printf("error closing file reader: %v", err)
		}
	}(reader)

	contentType := "application/octet-stream"
	if file.MimeType != nil {
		contentType = *file.MimeType
	}

	contentLength := int64(0)
	if file.Size != nil {
		contentLength = *file.Size
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", contentLength))

	http.ServeContent(w, r, file.Name, file.UpdatedAt, reader)
}

// Delete godoc
//
//	@summary	Delete a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		fileId	path	uint	true	"File identifier"
//	@success	204
//	@failure	404	{object}	handler.ErrorResponse
//	@failure	500	{object}	handler.ErrorResponse
//	@router		/api/v1/files/{fileId} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseFileId(r)
	if err != nil {
		writeInvalidFileIdError(w)
		return
	}

	err = h.service.Delete(r.Context(), id)
	if errors.Is(err, ErrFileNotFound) {
		writeFileNotFoundError(w)
		return
	}
	if err != nil {
		handler.WriteInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeInvalidFileIdError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusBadRequest, "invalid file id")
}

func writeFileNotFoundError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusNotFound, "file not found")
}

func writeFileAlreadyExistsError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusConflict, "file already exists")
}

func writeFileNotCompleteError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusNotFound, "file not complete")
}

func writeFileAlreadyCompleteError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusConflict, "file already complete")
}

func parseFileId(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "fileId"), 10, 64)
}

func parseVersionId(r *http.Request) (*int64, error) {
	if !r.URL.Query().Has("versionId") {
		return nil, nil
	}

	versionId, err := strconv.ParseInt(r.URL.Query().Get("versionId"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &versionId, nil
}

func writeInvalidVersionIdError(w http.ResponseWriter) {
	handler.WriteError(w, http.StatusBadRequest, "invalid version id")
}
