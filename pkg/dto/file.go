package dto

import (
	"app/pkg/model"
	"time"
)

type CreateFileDto struct {
	Name string `json:"name" binding:"required" example:"example.pdf"`
	Size int64  `json:"size" binding:"required" example:"1024"`
	Path string `json:"-" binding:"required" example:"/tmp/example.pdf"`
}

type FileResponseDto struct {
	ID        uint   `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name      string `json:"name" example:"example.pdf"`
	Size      int64  `json:"size" example:"1024"`
}

func ToFileResponse(file *model.File) *FileResponseDto {
	return &FileResponseDto{
		ID:        file.ID,
		CreatedAt: file.CreatedAt.Format(time.RFC3339),
		UpdatedAt: file.UpdatedAt.Format(time.RFC3339),
		Name:      file.Name,
		Size:      file.Size,
	}
}

type ListFilesResponseFileDto struct {
	ID        uint   `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name      string `json:"name" example:"example.pdf"`
	Size      int64  `json:"size" example:"1024"`
}

func ToListFilesResponseFile(file *model.File) *ListFilesResponseFileDto {
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

func ToListFilesResponse(files []model.File, total int64) *ListFilesResponseDto {
	listFilesResponseFileDtos := make([]ListFilesResponseFileDto, len(files))
	for i, file := range files {
		listFilesResponseFileDtos[i] = *ToListFilesResponseFile(&file)
	}

	return &ListFilesResponseDto{
		Files: listFilesResponseFileDtos,
		Total: total,
	}
}
