package file

import (
	"mime/multipart"
	"time"
)

type File struct {
	ID         int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Size       *int64
	Path       *string
	MimeType   *string
	IsComplete bool
}

type CreateFileRequest struct {
	Name string `json:"name" validate:"required" example:"example.pdf"`
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
	IsComplete bool    `json:"isComplete" example:"true"`
}

type ListFilesResponse struct {
	Files  []FileResponse `json:"files"`
	Limit  int64          `json:"limit"`
	Offset int64          `json:"offset"`
}

func (f File) ToResponse() FileResponse {
	return FileResponse{
		ID:         f.ID,
		CreatedAt:  f.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  f.UpdatedAt.Format(time.RFC3339),
		Name:       f.Name,
		Size:       f.Size,
		MimeType:   f.MimeType,
		IsComplete: f.IsComplete,
	}
}

func ToListResponse(files []File, limit, offset int64) ListFilesResponse {
	response := make([]FileResponse, len(files))
	for i, file := range files {
		response[i] = file.ToResponse()
	}
	return ListFilesResponse{
		Files:  response,
		Limit:  limit,
		Offset: offset,
	}
}
