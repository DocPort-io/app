package dto

import (
	"app/pkg/database"
	"net/http"
	"time"
)

type PaginationParams struct {
	Limit  int64
	Offset int64
}

type FindAllProjectsParams struct {
	*PaginationParams
}

type FindAllProjectsResult struct {
	Projects []*database.ListProjectsWithLocationsRow
	Total    int64
	Limit    int64
	Offset   int64
}

type FindProjectByIdParams struct {
	ID int64
}

type FindProjectByIdResult struct {
	Project *database.Project
}

type CreateProjectParams struct {
	Slug string
	Name string
}

type CreateProjectResult struct {
	Project *database.Project
}

type UpdateProjectParams struct {
	Slug string
	Name string
	ID   int64
}

type UpdateProjectResult struct {
	Project *database.Project
}

type DeleteProjectParams struct {
	ID int64
}

type CreateProjectDto struct {
	Slug string `json:"slug" binding:"required" example:"project-x"`
	Name string `json:"name" binding:"required" example:"Project X"`
}

func (c *CreateProjectDto) Bind(r *http.Request) error {
	return nil
}

type UpdateProjectDto struct {
	Slug string `json:"slug" example:"project-x"`
	Name string `json:"name" example:"Project X"`
}

func (u *UpdateProjectDto) Bind(r *http.Request) error {
	return nil
}

type ProjectResponseDto struct {
	ID        int64  `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Slug      string `json:"slug" example:"project-x"`
	Name      string `json:"name" example:"Project X"`
}

func (p *ProjectResponseDto) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToProjectResponse(project *database.Project) *ProjectResponseDto {
	return &ProjectResponseDto{
		ID:        project.ID,
		CreatedAt: project.CreatedAt.Format(time.RFC3339),
		UpdatedAt: project.UpdatedAt.Format(time.RFC3339),
		Slug:      project.Slug,
		Name:      project.Name,
	}
}

type ListProjectsResponseProjectLocationDto struct {
	Name *string  `json:"name" example:"Office"`
	Lat  *float64 `json:"lat" example:"52.520008"`
	Lon  *float64 `json:"lon" example:"13.404954"`
}

func ToListProjectsResponseProjectLocationDto(project *database.ListProjectsWithLocationsRow) *ListProjectsResponseProjectLocationDto {
	if project.LocationName == nil {
		return nil
	}

	return &ListProjectsResponseProjectLocationDto{
		Name: project.LocationName,
		Lat:  project.LocationLat,
		Lon:  project.LocationLon,
	}
}

type ListProjectsResponseProjectDto struct {
	ID        int64                                   `json:"id" example:"1"`
	CreatedAt string                                  `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string                                  `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Slug      string                                  `json:"slug" example:"project-x"`
	Name      string                                  `json:"name" example:"Project X"`
	Location  *ListProjectsResponseProjectLocationDto `json:"location"`
}

func ToListProjectsResponseProject(project *database.ListProjectsWithLocationsRow) *ListProjectsResponseProjectDto {
	return &ListProjectsResponseProjectDto{
		ID:        project.ID,
		CreatedAt: project.CreatedAt.Format(time.RFC3339),
		UpdatedAt: project.UpdatedAt.Format(time.RFC3339),
		Slug:      project.Slug,
		Name:      project.Name,
		Location:  ToListProjectsResponseProjectLocationDto(project),
	}
}

type ListProjectsResponseDto struct {
	Projects []ListProjectsResponseProjectDto `json:"projects"`
	Total    int64                            `json:"total" example:"1"`
	Limit    int64                            `json:"limit" example:"100"`
	Offset   int64                            `json:"offset" example:"0"`
}

func (l *ListProjectsResponseDto) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToListProjectsResponse(result *FindAllProjectsResult) *ListProjectsResponseDto {
	projectResponseProjectDtos := make([]ListProjectsResponseProjectDto, len(result.Projects))
	for i, project := range result.Projects {
		projectResponseProjectDtos[i] = *ToListProjectsResponseProject(project)
	}
	return &ListProjectsResponseDto{
		Projects: projectResponseProjectDtos,
		Total:    result.Total,
		Limit:    result.Limit,
		Offset:   result.Offset,
	}
}
