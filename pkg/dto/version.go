package dto

import (
	"app/pkg/model"
	"app/pkg/util"
	"time"
)

type CreateVersionDto struct {
	Name        string `json:"name" binding:"required" example:"v0.0.1"`
	Description string `json:"description" example:"First version of the project"`
	ProjectId   uint   `json:"projectId" example:"1"`
}

func (dto *CreateVersionDto) ToModel() *model.Version {
	return &model.Version{
		Name:        dto.Name,
		Description: dto.Description,
		ProjectId:   dto.ProjectId,
	}
}

type UpdateVersionDto struct {
	Name        string `json:"name" example:"v0.0.1"`
	Description string `json:"description" example:"First version of the project"`
}

func (dto *UpdateVersionDto) ToModel() *model.Version {
	return &model.Version{
		Name:        dto.Name,
		Description: dto.Description,
	}
}

type UploadFileToVersionDto struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}

func ToUploadFileToVersionDto(file *util.MultipartFile) *UploadFileToVersionDto {
	return &UploadFileToVersionDto{
		Name: file.Name,
		Size: file.Size,
		Path: file.Path,
	}
}

type VersionResponseDto struct {
	ID          uint   `json:"id" example:"1"`
	CreatedAt   string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt   string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name        string `json:"name" example:"v0.0.1"`
	Description string `json:"description" example:"First version of the project"`
	ProjectId   uint   `json:"projectId" example:"1"`
}

func ToVersionResponse(version *model.Version) *VersionResponseDto {
	return &VersionResponseDto{
		ID:          version.ID,
		CreatedAt:   version.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   version.UpdatedAt.Format(time.RFC3339),
		Name:        version.Name,
		Description: version.Description,
		ProjectId:   version.ProjectId,
	}
}

type ListVersionsResponseVersionDto struct {
	ID        uint   `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name      string `json:"name" example:"First version of the project"`
	ProjectId uint   `json:"projectId" example:"1"`
}

func ToListVersionsResponseVersion(version *model.Version) *ListVersionsResponseVersionDto {
	return &ListVersionsResponseVersionDto{
		ID:        version.ID,
		CreatedAt: version.CreatedAt.Format(time.RFC3339),
		UpdatedAt: version.UpdatedAt.Format(time.RFC3339),
		Name:      version.Name,
		ProjectId: version.ProjectId,
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
