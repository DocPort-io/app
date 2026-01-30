package file

import (
	"app/pkg/database"
	"app/pkg/paginate"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

// Service layer

type FindAllFilesParams struct {
	VersionID int64
	*paginate.Pagination
}

type FindAllFilesResult struct {
	Files  []*database.File
	Total  int64
	Limit  int64
	Offset int64
}

type FindFileByIdParams struct {
	ID int64
}

type FindFileByIdResult struct {
	File *database.File
}

type CreateFileParams struct {
	Name string
}

type CreateFileResult struct {
	File *database.File
}

type UploadFileParams struct {
	ID         int64
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type UploadFileResult struct {
	File *database.File
}

type DownloadFileParams struct {
	ID int64
}

type DownloadFileResult struct {
	File   *database.File
	Reader io.ReadSeekCloser
}

type DeleteFileParams struct {
	ID int64
}

// Controller layer

type CreateFileRequest struct {
	Name string `json:"name" validate:"required" example:"example.pdf"`
}

func (c *CreateFileRequest) Bind(r *http.Request) error {
	v := validator.New(validator.WithRequiredStructEnabled())
	return v.Struct(c)
}

type UploadFileRequest struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type FileResponse struct {
	ID         int64   `json:"id" example:"1"`
	CreatedAt  string  `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt  string  `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name       string  `json:"name" example:"example.pdf"`
	Size       *int64  `json:"size" example:"1024"`
	MimeType   *string `json:"mimeType" example:"application/pdf"`
	IsComplete bool    `json:"isComplete" example:"false"`
}

func (f *FileResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToFileResponse(file *database.File) *FileResponse {
	return &FileResponse{
		ID:         file.ID,
		CreatedAt:  file.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  file.UpdatedAt.Time.Format(time.RFC3339),
		Name:       file.Name,
		Size:       file.Size,
		MimeType:   file.MimeType,
		IsComplete: file.IsComplete,
	}
}

type ListFilesResponseFile struct {
	ID         int64   `json:"id" example:"1"`
	CreatedAt  string  `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt  string  `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name       string  `json:"name" example:"example.pdf"`
	Size       *int64  `json:"size" example:"1024"`
	MimeType   *string `json:"mimeType" example:"application/pdf"`
	IsComplete bool    `json:"isComplete" example:"false"`
}

func ToListFilesResponseFile(file *database.File) *ListFilesResponseFile {
	return &ListFilesResponseFile{
		ID:         file.ID,
		CreatedAt:  file.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  file.UpdatedAt.Time.Format(time.RFC3339),
		Name:       file.Name,
		Size:       file.Size,
		MimeType:   file.MimeType,
		IsComplete: file.IsComplete,
	}
}

type ListFilesResponse struct {
	Files  []ListFilesResponseFile `json:"files"`
	Total  int64                   `json:"total" example:"1"`
	Limit  int64                   `json:"limit" example:"100"`
	Offset int64                   `json:"offset" example:"0"`
}

func (l *ListFilesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToListFilesResponse(result *FindAllFilesResult) *ListFilesResponse {
	listFilesResponseFileDtos := make([]ListFilesResponseFile, len(result.Files))
	for i, file := range result.Files {
		listFilesResponseFileDtos[i] = *ToListFilesResponseFile(file)
	}
	return &ListFilesResponse{
		Files:  listFilesResponseFileDtos,
		Total:  result.Total,
		Limit:  result.Limit,
		Offset: result.Offset,
	}
}
