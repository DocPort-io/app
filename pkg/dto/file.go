package dto

import (
	"app/pkg/database"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

// Service layer

type FindAllFilesParams struct {
	VersionID int64
}

type FindAllFilesResult struct {
	Files []*database.File
	Total int64
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
	ID         int64   `json:"id" example:"1"`
	CreatedAt  string  `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt  string  `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name       string  `json:"name" example:"example.pdf"`
	Size       *int64  `json:"size" example:"1024"`
	MimeType   *string `json:"mimeType" example:"application/pdf"`
	IsComplete bool    `json:"isComplete" example:"false"`
}

func (f *FileResponseDto) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToFileResponse(file *database.File) *FileResponseDto {
	return &FileResponseDto{
		ID:         file.ID,
		CreatedAt:  file.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  file.UpdatedAt.Format(time.RFC3339),
		Name:       file.Name,
		Size:       file.Size,
		MimeType:   file.MimeType,
		IsComplete: file.IsComplete,
	}
}

type ListFilesResponseFileDto struct {
	ID         int64   `json:"id" example:"1"`
	CreatedAt  string  `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt  string  `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name       string  `json:"name" example:"example.pdf"`
	Size       *int64  `json:"size" example:"1024"`
	MimeType   *string `json:"mimeType" example:"application/pdf"`
	IsComplete bool    `json:"isComplete" example:"false"`
}

func ToListFilesResponseFile(file *database.File) *ListFilesResponseFileDto {
	return &ListFilesResponseFileDto{
		ID:         file.ID,
		CreatedAt:  file.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  file.UpdatedAt.Format(time.RFC3339),
		Name:       file.Name,
		Size:       file.Size,
		MimeType:   file.MimeType,
		IsComplete: file.IsComplete,
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
