package controller

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/httputil"
	"app/pkg/service"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

const FileCtxKey = "file"

type FileController struct {
	fileService service.FileService
}

func NewFileController(fileService service.FileService) *FileController {
	return &FileController{fileService: fileService}
}

func (c *FileController) FileCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileId, err := httputil.URLParamInt64(r, "fileId")
		if err != nil {
			httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
			return
		}

		file, err := c.fileService.FindFileById(r.Context(), fileId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNotFound) {
				httputil.Render(w, r, apperrors.ErrHTTPNotFoundError())
				return
			}

			httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
			return
		}

		ctx := context.WithValue(r.Context(), FileCtxKey, file)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getFile(ctx context.Context) *database.File {
	return ctx.Value(FileCtxKey).(*database.File)
}

// FindAllFiles godoc
//
//	@summary	Find all files
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		versionId	query		uint	false	"Version identifier"
//	@success	200			{object}	dto.ListFilesResponseDto
//	@failure	400			{object}	apperrors.ErrResponse
//	@failure	500			{object}	apperrors.ErrResponse
//	@router		/api/v1/files [get]
func (c *FileController) FindAllFiles(w http.ResponseWriter, r *http.Request) {
	versionId, err := httputil.QueryParamInt64(r, "versionId", true)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	files, total, err := c.fileService.FindAllFiles(r.Context(), versionId)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToListFilesResponse(files, total))
}

// GetFile godoc
//
//	@summary	Get a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		fileId	path		uint	true	"File identifier"
//	@success	200		{object}	dto.FileResponseDto
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	404		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/files/{fileId} [get]
func (c *FileController) GetFile(w http.ResponseWriter, r *http.Request) {
	file := getFile(r.Context())
	httputil.Render(w, r, dto.ToFileResponse(file))
}

// CreateFile godoc
//
//	@summary	Create a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		request	body		dto.CreateFileDto	true	"Create a file"
//	@success	201		{object}	dto.FileResponseDto
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/files [post]
func (c *FileController) CreateFile(w http.ResponseWriter, r *http.Request) {
	input := &dto.CreateFileDto{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	file, err := c.fileService.CreateFile(r.Context(), input)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	httputil.Render(w, r, dto.ToFileResponse(file))
}

// UploadFile godoc
//
//	@summary	Upload a file
//	@tags		files
//	@accept		multipart/form-data
//	@produce	json
//	@param		fileId	path		uint	true	"File identifier"
//	@param		file	formData	file	true	"File to upload"
//	@success	201		{object}	dto.FileResponseDto
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	404		{object}	apperrors.ErrResponse
//	@failure	409		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/files/{fileId}/upload [post]
func (c *FileController) UploadFile(w http.ResponseWriter, r *http.Request) {
	file := getFile(r.Context())

	r.Body = http.MaxBytesReader(w, r.Body, 1_073_741_824)

	multipartFile, multipartFileHeader, err := r.FormFile("file")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	uploadFileDto := &dto.UploadFileDto{File: multipartFile, FileHeader: multipartFileHeader}

	file, err = c.fileService.UploadFile(r.Context(), file.ID, uploadFileDto)
	if err != nil {
		if errors.Is(err, service.ErrFileAlreadyExists) {
			httputil.Render(w, r, apperrors.ErrHTTPConflictError(err))
			return
		}

		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	httputil.Render(w, r, dto.ToFileResponse(file))
}

// DownloadFile godoc
//
//	@summary	Download a file
//	@tags		files
//	@accept		json
//	@param		fileId	path	uint	true	"File identifier"
//	@success	200
//	@failure	400	{object}	apperrors.ErrResponse
//	@failure	404	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/files/{fileId}/download [get]
func (c *FileController) DownloadFile(w http.ResponseWriter, r *http.Request) {
	file := getFile(r.Context())

	file, reader, err := c.fileService.DownloadFile(r.Context(), file.ID)
	if err != nil {
		if errors.Is(err, service.ErrIncompleteFile) {
			httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
			return
		}

		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
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

// DeleteFile godoc
//
//	@summary	Delete a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		fileId	path	uint	true	"File identifier"
//	@success	204
//	@failure	404	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/files/{fileId} [delete]
func (c *FileController) DeleteFile(w http.ResponseWriter, r *http.Request) {
	file := getFile(r.Context())

	err := c.fileService.DeleteFile(r.Context(), file.ID)
	if err != nil {
		log.Printf("error deleting file: %v", err)
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}
