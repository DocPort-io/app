package dto

import (
	"app/pkg/database"
	"app/pkg/paginate"
	"net/http"
	"time"
)

// Service layer

type FindAllVersionsParams struct {
	ProjectID int64
	*paginate.Pagination
}

type FindAllVersionsResult struct {
	Versions []*database.Version
	Total    int64
	Limit    int64
	Offset   int64
}

type FindVersionByIdParams struct {
	ID int64
}

type FindVersionByIdResult struct {
	Version *database.Version
}

type CreateVersionParams struct {
	Name        string
	Description *string
	ProjectID   int64
}

type CreateVersionResult struct {
	Version *database.Version
}

type UpdateVersionParams struct {
	ID          int64
	Name        string
	Description *string
}

type UpdateVersionResult struct {
	Version *database.Version
}

type DeleteVersionParams struct {
	ID int64
}

type AttachFileToVersionParams struct {
	VersionID int64
	FileID    int64
}

type DetachFileFromVersionParams struct {
	VersionID int64
	FileID    int64
}

// Controller layer

type CreateVersionRequest struct {
	Name        string  `json:"name" binding:"required" example:"v0.0.1"`
	Description *string `json:"description" example:"First version of the project"`
	ProjectId   int64   `json:"projectId" example:"1"`
}

func (v *CreateVersionRequest) Bind(r *http.Request) error {
	return nil
}

type UpdateVersionRequest struct {
	Name        string  `json:"name" example:"v0.0.1"`
	Description *string `json:"description" example:"First version of the project"`
}

func (v *UpdateVersionRequest) Bind(r *http.Request) error {
	return nil
}

type AttachFileToVersionRequest struct {
	FileId int64 `json:"fileId" example:"1"`
}

func (v *AttachFileToVersionRequest) Bind(r *http.Request) error {
	return nil
}

type DetachFileFromVersionRequest struct {
	FileId int64 `json:"fileId" example:"1"`
}

func (v *DetachFileFromVersionRequest) Bind(r *http.Request) error {
	return nil
}

type VersionResponse struct {
	ID          int64   `json:"id" example:"1"`
	CreatedAt   string  `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt   string  `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name        string  `json:"name" example:"v0.0.1"`
	Description *string `json:"description" example:"First version of the project"`
	ProjectId   int64   `json:"projectId" example:"1"`
}

func (v *VersionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToVersionResponse(version *database.Version) *VersionResponse {
	return &VersionResponse{
		ID:          version.ID,
		CreatedAt:   version.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   version.UpdatedAt.Format(time.RFC3339),
		Name:        version.Name,
		Description: version.Description,
		ProjectId:   version.ProjectID,
	}
}

type listVersionsResponseVersion struct {
	ID        int64  `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name      string `json:"name" example:"First version of the project"`
	ProjectId int64  `json:"projectId" example:"1"`
}

func toListVersionsResponseVersion(version *database.Version) *listVersionsResponseVersion {
	return &listVersionsResponseVersion{
		ID:        version.ID,
		CreatedAt: version.CreatedAt.Format(time.RFC3339),
		UpdatedAt: version.UpdatedAt.Format(time.RFC3339),
		Name:      version.Name,
		ProjectId: version.ProjectID,
	}
}

type ListVersionsResponse struct {
	Versions []listVersionsResponseVersion `json:"versions"`
	Total    int64                         `json:"total" example:"1"`
	Limit    int64                         `json:"limit" example:"100"`
	Offset   int64                         `json:"offset" example:"0"`
}

func (l *ListVersionsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToListVersionsResponse(result *FindAllVersionsResult) *ListVersionsResponse {
	versionsResponseVersionDtos := make([]listVersionsResponseVersion, len(result.Versions))
	for i, version := range result.Versions {
		versionsResponseVersionDtos[i] = *toListVersionsResponseVersion(version)
	}
	return &ListVersionsResponse{
		Versions: versionsResponseVersionDtos,
		Total:    result.Total,
		Limit:    result.Limit,
		Offset:   result.Offset,
	}
}
