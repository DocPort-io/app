package dto

import (
	"app/pkg/model"
	"time"
)

type CreateVersionDto struct {
	Name        string `json:"name" binding:"required" example:"v0.0.1"`
	Description string `json:"description" example:"First version of the project"`
	ProjectID   uint   `json:"projectId" example:"1"`
}

func (dto *CreateVersionDto) ToModel() *model.Version {
	return &model.Version{
		Name:        dto.Name,
		Description: dto.Description,
		ProjectID:   dto.ProjectID,
	}
}

type VersionResponseDto struct {
	ID          uint   `json:"id" example:"1"`
	CreatedAt   string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt   string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name        string `json:"name" example:"v0.0.1"`
	Description string `json:"description" example:"First version of the project"`
	ProjectID   uint   `json:"projectId" example:"1"`
}

func ToVersionResponse(version *model.Version) *VersionResponseDto {
	return &VersionResponseDto{
		ID:          version.ID,
		CreatedAt:   version.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   version.UpdatedAt.Format(time.RFC3339),
		Name:        version.Name,
		Description: version.Description,
	}
}

type ListVersionsResponseVersionDto struct {
	ID        uint   `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name      string `json:"name" example:"First version of the project"`
}

func ToListVersionsResponseVersion(version *model.Version) *ListVersionsResponseVersionDto {
	return &ListVersionsResponseVersionDto{
		ID:        version.ID,
		CreatedAt: version.CreatedAt.Format(time.RFC3339),
		UpdatedAt: version.UpdatedAt.Format(time.RFC3339),
		Name:      version.Name,
	}
}

type ListVersionsResponseDto struct {
	Versions []ListVersionsResponseVersionDto `json:"versions"`
	Total    int64                            `json:"total" example:"1"`
}

func ToListVersionsResponse(versions []model.Version, total int64) *ListVersionsResponseDto {
	versionsResponseVersionDtos := make([]ListVersionsResponseVersionDto, len(versions))
	for i, version := range versions {
		versionsResponseVersionDtos[i] = *ToListVersionsResponseVersion(&version)
	}
	return &ListVersionsResponseDto{
		Versions: versionsResponseVersionDtos,
		Total:    total,
	}
}
