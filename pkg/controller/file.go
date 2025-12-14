package controller

import (
	"app/pkg/apperrors"
	"app/pkg/dto"
	"app/pkg/httputil"
	"app/pkg/service"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

type FileController struct {
	fileService *service.FileService
}

func NewFileController(fileService *service.FileService) *FileController {
	return &FileController{fileService: fileService}
}

// FindAllFiles godoc
//
//	@summary	Find all files
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		versionId	query		uint	false	"Version identifier"
//	@success	200			{object}	dto.ListFilesResponseDto
//	@router		/files [get]
func (c *FileController) FindAllFiles(w http.ResponseWriter, r *http.Request) {
	versionId, err := httputil.QueryParamInt64(r, "versionId", true)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	files, total, err := c.fileService.FindAllFiles(r.Context(), versionId)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrInternalServerError(err))
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
//	@param		id	path		uint	true	"File identifier"
//	@success	200	{object}	dto.FileResponseDto
//	@router		/files/{id} [get]
func (c *FileController) GetFile(w http.ResponseWriter, r *http.Request) {
	fileId, err := httputil.URLParamInt64(r, "fileId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	file, err := c.fileService.FindFileById(r.Context(), fileId)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}

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
//	@router		/files [post]
func (c *FileController) CreateFile(w http.ResponseWriter, r *http.Request) {
	input := &dto.CreateFileDto{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	file, err := c.fileService.CreateFile(r.Context(), input)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToFileResponse(file))
}

// UploadFile godoc
//
//	@summary	Upload a file
//	@tags		files
//	@accept		multipart/form-data
//	@produce	json
//	@param		id		path		uint	true	"File identifier"
//	@param		file	formData	file	true	"File to upload"
//	@success	201		{object}	dto.FileResponseDto
//	@failure	400		{object}	apperrors.ErrResponse
//	@router		/files/{id}/upload [post]
func (c *FileController) UploadFile(w http.ResponseWriter, r *http.Request) {
	fileId, err := httputil.URLParamInt64(r, "fileId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	multipartFile, multipartFileHeader, err := r.FormFile("file")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	uploadFileDto := &dto.UploadFileDto{File: multipartFile, FileHeader: multipartFileHeader}

	file, err := c.fileService.UploadFile(r.Context(), fileId, uploadFileDto)
	if err != nil {
		if errors.Is(err, service.ErrFileAlreadyExists) {
			httputil.Render(w, r, apperrors.ErrBadRequestError(err))
			return
		}

		httputil.Render(w, r, apperrors.ErrInternalServerError(err))
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
//	@param		id	path	uint	true	"File identifier"
//	@success	200
//	@failure	400	{object}	apperrors.ErrResponse
//	@router		/files/{id}/download [get]
func (c *FileController) DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileId, err := httputil.URLParamInt64(r, "fileId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	file, reader, err := c.fileService.DownloadFile(r.Context(), fileId)
	if err != nil {
		if errors.Is(err, service.ErrIncompleteFile) {
			httputil.Render(w, r, apperrors.ErrBadRequestError(err))
			return
		}

		httputil.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			log.Printf("error closing file reader: %v", err)
		}
	}(reader)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	w.Header().Set("Content-Type", *file.MimeType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", *file.Size))

	_, err = io.Copy(w, reader)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}
}

// DeleteFile godoc
//
//	@summary	Delete a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		id	path	uint	true	"File identifier"
//	@success	204
//	@router		/files/{id} [delete]
func (c *FileController) DeleteFile(w http.ResponseWriter, r *http.Request) {
	fileId, err := httputil.URLParamInt64(r, "fileId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	err = c.fileService.DeleteFile(r.Context(), fileId)
	if err != nil {
		log.Printf("error deleting file: %v", err)
		httputil.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}
