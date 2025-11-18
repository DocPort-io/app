package dto

import (
	"app/pkg/model"
	"time"
)

type VersionResponseDto struct {
	ID          uint   `json:"id" example:"1"`
	CreatedAt   string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt   string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name        string `json:"name" example:"v0.0.1"`
	Description string `json:"description" example:"First version of the project"`
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
