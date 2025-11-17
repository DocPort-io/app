package dto

import "app/pkg/model"

type VersionResponseDto struct {
	ID uint `json:"id" example:"1"`
}

func ToVersionResponse(version *model.Version) *VersionResponseDto {
	return &VersionResponseDto{
		ID: version.ID,
	}
}
