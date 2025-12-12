package dto

import (
	"app/pkg/database"
	"net/http"
	"time"
)

type CreateVersionDto struct {
	Name        string  `json:"name" binding:"required" example:"v0.0.1"`
	Description *string `json:"description" example:"First version of the project"`
	ProjectId   int64   `json:"projectId" example:"1"`
}

func (v *CreateVersionDto) Bind(r *http.Request) error {
	return nil
}

type UpdateVersionDto struct {
	Name        string  `json:"name" example:"v0.0.1"`
	Description *string `json:"description" example:"First version of the project"`
}

func (v *UpdateVersionDto) Bind(r *http.Request) error {
	return nil
}

type AttachFileToVersionDto struct {
	FileId int64 `json:"fileId" example:"1"`
}

func (v *AttachFileToVersionDto) Bind(r *http.Request) error {
	return nil
}

type DetachFileFromVersionDto struct {
	FileId int64 `json:"fileId" example:"1"`
}

func (v *DetachFileFromVersionDto) Bind(r *http.Request) error {
	return nil
}

type VersionResponseDto struct {
	ID          int64   `json:"id" example:"1"`
	CreatedAt   string  `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt   string  `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name        string  `json:"name" example:"v0.0.1"`
	Description *string `json:"description" example:"First version of the project"`
	ProjectId   int64   `json:"projectId" example:"1"`
}

func (v *VersionResponseDto) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToVersionResponse(version *database.Version) *VersionResponseDto {
	return &VersionResponseDto{
		ID:          version.ID,
		CreatedAt:   version.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   version.UpdatedAt.Format(time.RFC3339),
		Name:        version.Name,
		Description: version.Description,
		ProjectId:   version.ProjectID,
	}
}

type ListVersionsResponseVersionDto struct {
	ID        int64  `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name      string `json:"name" example:"First version of the project"`
	ProjectId int64  `json:"projectId" example:"1"`
}

func ToListVersionsResponseVersion(version *database.Version) *ListVersionsResponseVersionDto {
	return &ListVersionsResponseVersionDto{
		ID:        version.ID,
		CreatedAt: version.CreatedAt.Format(time.RFC3339),
		UpdatedAt: version.UpdatedAt.Format(time.RFC3339),
		Name:      version.Name,
		ProjectId: version.ProjectID,
	}
}

type ListVersionsResponseDto struct {
	Versions []ListVersionsResponseVersionDto `json:"versions"`
	Total    int64                            `json:"total" example:"1"`
}

func (l *ListVersionsResponseDto) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToListVersionsResponse(versions []*database.Version, total int64) *ListVersionsResponseDto {
	versionsResponseVersionDtos := make([]ListVersionsResponseVersionDto, len(versions))
	for i, version := range versions {
		versionsResponseVersionDtos[i] = *ToListVersionsResponseVersion(version)
	}
	return &ListVersionsResponseDto{
		Versions: versionsResponseVersionDtos,
		Total:    total,
	}
}
