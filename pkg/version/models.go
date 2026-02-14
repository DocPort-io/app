package version

import "time"

type Version struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description *string
	ProjectID   int64
}

type CreateVersionRequest struct {
	Name        string  `json:"name" validate:"required" example:"v0.0.1"`
	Description *string `json:"description,omitempty" validate:"omitempty" example:"First version of the project"`
	ProjectId   int64   `json:"projectId" validate:"required,gt=0" example:"1"`
}

type UpdateVersionRequest struct {
	Name        string  `json:"name" validate:"required" example:"v0.0.1"`
	Description *string `json:"description,omitempty" validate:"omitempty" example:"First version of the project"`
}

type AttachFileRequest struct {
	FileID int64 `json:"fileId" validate:"required,gt=0" example:"1"`
}

type DetachFileRequest struct {
	FileID int64 `json:"fileId" validate:"required,gt=0" example:"1"`
}

type VersionResponse struct {
	ID          int64   `json:"id" example:"1"`
	CreatedAt   string  `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt   string  `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name        string  `json:"name" example:"v0.0.1"`
	Description *string `json:"description,omitempty" example:"First version of the project"`
	ProjectID   int64   `json:"projectId" example:"1"`
}

type ListVersionsResponse struct {
	Versions []VersionResponse `json:"versions"`
	Limit    int64             `json:"limit"`
	Offset   int64             `json:"offset"`
}

func (v Version) ToResponse() VersionResponse {
	return VersionResponse{
		ID:          v.ID,
		CreatedAt:   v.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
		Name:        v.Name,
		Description: v.Description,
		ProjectID:   v.ProjectID,
	}
}

func ToListResponse(versions []Version, limit, offset int64) ListVersionsResponse {
	response := make([]VersionResponse, len(versions))
	for i, version := range versions {
		response[i] = version.ToResponse()
	}
	return ListVersionsResponse{
		Versions: response,
		Limit:    limit,
		Offset:   offset,
	}
}
