package project

import "time"

type Project struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Slug      string
	Name      string
}

type CreateProjectRequest struct {
	Slug string `json:"slug" validate:"required" example:"project-x"`
	Name string `json:"name" validate:"required" example:"Project X"`
}

type UpdateProjectRequest struct {
	Slug string `json:"slug" validate:"required" example:"project-x"`
	Name string `json:"name" validate:"required" example:"Project X"`
}

type ProjectResponse struct {
	ID        int64  `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Slug      string `json:"slug" example:"project-x"`
	Name      string `json:"name" example:"Project X"`
}

type ListProjectsResponse struct {
	Projects []ProjectResponse `json:"projects"`
	Limit    int64             `json:"limit"`
	Offset   int64             `json:"offset"`
}

func (p Project) ToResponse() ProjectResponse {
	return ProjectResponse{
		ID:        p.ID,
		CreatedAt: p.CreatedAt.Format(time.RFC3339),
		UpdatedAt: p.UpdatedAt.Format(time.RFC3339),
		Slug:      p.Slug,
		Name:      p.Name,
	}
}

func ToListResponse(projects []Project, limit, offset int64) ListProjectsResponse {
	response := make([]ProjectResponse, len(projects))
	for i, project := range projects {
		response[i] = project.ToResponse()
	}
	return ListProjectsResponse{
		Projects: response,
		Limit:    limit,
		Offset:   offset,
	}
}
