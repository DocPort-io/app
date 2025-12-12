package dto

import (
	"app/pkg/database"
	"mime/multipart"
	"net/http"
	"time"
)

type CreateFileDto struct {
	Name string `json:"name" binding:"required" example:"example.pdf"`
}

func (c *CreateFileDto) Bind(r *http.Request) error {
	return nil
}

type UploadFileDto struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type FileResponseDto struct {
	ID        int64  `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name      string `json:"name" example:"example.pdf"`
	Size      *int64 `json:"size" example:"1024"`
}

func (f *FileResponseDto) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToFileResponse(file *database.File) *FileResponseDto {
	return &FileResponseDto{
		ID:        file.ID,
		CreatedAt: file.CreatedAt.Format(time.RFC3339),
		UpdatedAt: file.UpdatedAt.Format(time.RFC3339),
		Name:      file.Name,
		Size:      file.Size,
	}
}

type ListFilesResponseFileDto struct {
	ID        int64  `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name      string `json:"name" example:"example.pdf"`
	Size      *int64 `json:"size" example:"1024"`
}

func ToListFilesResponseFile(file *database.File) *ListFilesResponseFileDto {
	return &ListFilesResponseFileDto{
		ID:        file.ID,
		CreatedAt: file.CreatedAt.Format(time.RFC3339),
		UpdatedAt: file.UpdatedAt.Format(time.RFC3339),
		Name:      file.Name,
		Size:      file.Size,
	}
}

type ListFilesResponseDto struct {
	Files []ListFilesResponseFileDto `json:"files"`
	Total int64                      `json:"total" example:"1"`
}

func (l *ListFilesResponseDto) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToListFilesResponse(files []*database.File, total int64) *ListFilesResponseDto {
	listFilesResponseFileDtos := make([]ListFilesResponseFileDto, len(files))
	for i, file := range files {
		listFilesResponseFileDtos[i] = *ToListFilesResponseFile(file)
	}

	return &ListFilesResponseDto{
		Files: listFilesResponseFileDtos,
		Total: total,
	}
}
