package controller

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/httputil"
	"app/pkg/paginate"
	"app/pkg/service"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/dustin/go-humanize"
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

		findRes, err := c.fileService.FindFileById(r.Context(), &dto.FindFileByIdParams{ID: fileId})
		if err != nil {
			if errors.Is(err, apperrors.ErrNotFound) {
				httputil.Render(w, r, apperrors.ErrHTTPNotFoundError())
				return
			}

			httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
			return
		}

		ctx := context.WithValue(r.Context(), FileCtxKey, findRes.File)
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
//	@param		limit		query		uint	false	"Max items per page (1-100)"
//	@param		offset		query		uint	false	"Items to skip before starting to collect the result set"
//	@success	200			{object}	dto.ListFilesResponse
//	@failure	400			{object}	apperrors.ErrResponse
//	@failure	500			{object}	apperrors.ErrResponse
//	@router		/api/v1/files [get]
func (c *FileController) FindAllFiles(w http.ResponseWriter, r *http.Request) {
	versionId, err := httputil.QueryParamInt64(r, "versionId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	pagination := paginate.GetPagination(r.Context())

	result, err := c.fileService.FindAllFiles(r.Context(), &dto.FindAllFilesParams{VersionID: versionId, Pagination: pagination})
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToListFilesResponse(result))
}

// GetFile godoc
//
//	@summary	Get a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		fileId	path		uint	true	"File identifier"
//	@success	200		{object}	dto.FileResponse
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
//	@param		request	body		dto.CreateFileRequest	true	"Create a file"
//	@success	201		{object}	dto.FileResponse
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/files [post]
func (c *FileController) CreateFile(w http.ResponseWriter, r *http.Request) {
	input := &dto.CreateFileRequest{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	createParams := &dto.CreateFileParams{Name: input.Name}
	createResult, err := c.fileService.CreateFile(r.Context(), createParams)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	httputil.Render(w, r, dto.ToFileResponse(createResult.File))
}

// UploadFile godoc
//
//	@summary	Upload a file
//	@tags		files
//	@accept		multipart/form-data
//	@produce	json
//	@param		fileId	path		uint	true	"File identifier"
//	@param		file	formData	file	true	"File to upload"
//	@success	201		{object}	dto.FileResponse
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	404		{object}	apperrors.ErrResponse
//	@failure	409		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/files/{fileId}/upload [post]
func (c *FileController) UploadFile(w http.ResponseWriter, r *http.Request) {
	file := getFile(r.Context())

	r.Body = http.MaxBytesReader(w, r.Body, 1*humanize.GiByte)

	multipartFile, multipartFileHeader, err := r.FormFile("file")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	uploadParams := &dto.UploadFileParams{ID: file.ID, File: multipartFile, FileHeader: multipartFileHeader}
	uploadResult, err := c.fileService.UploadFile(r.Context(), uploadParams)
	if err != nil {
		if errors.Is(err, service.ErrFileAlreadyExists) {
			httputil.Render(w, r, apperrors.ErrHTTPConflictError(err))
			return
		}

		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	httputil.Render(w, r, dto.ToFileResponse(uploadResult.File))
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

	downloadResult, err := c.fileService.DownloadFile(r.Context(), &dto.DownloadFileParams{ID: file.ID})
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
	}(downloadResult.Reader)

	contentType := "application/octet-stream"
	if downloadResult.File.MimeType != nil {
		contentType = *downloadResult.File.MimeType
	}

	contentLength := int64(0)
	if downloadResult.File.Size != nil {
		contentLength = *downloadResult.File.Size
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", downloadResult.File.Name))
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", contentLength))

	http.ServeContent(w, r, downloadResult.File.Name, downloadResult.File.UpdatedAt, downloadResult.Reader)
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

	err := c.fileService.DeleteFile(r.Context(), &dto.DeleteFileParams{ID: file.ID})
	if err != nil {
		log.Printf("error deleting file: %v", err)
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}
