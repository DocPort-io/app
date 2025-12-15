package dto

import (
	"app/pkg/database"
	"app/pkg/paginate"
	"net/http"
	"time"
)

// Service layer

type FindAllProjectsParams struct {
	*paginate.Pagination
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

// Controller layer

type CreateProjectRequest struct {
	Slug string `json:"slug" binding:"required" example:"project-x"`
	Name string `json:"name" binding:"required" example:"Project X"`
}

func (c *CreateProjectRequest) Bind(r *http.Request) error {
	return nil
}

type UpdateProjectRequest struct {
	Slug string `json:"slug" example:"project-x"`
	Name string `json:"name" example:"Project X"`
}

func (u *UpdateProjectRequest) Bind(r *http.Request) error {
	return nil
}

type ProjectResponse struct {
	ID        int64  `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Slug      string `json:"slug" example:"project-x"`
	Name      string `json:"name" example:"Project X"`
}

func (p *ProjectResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToProjectResponse(project *database.Project) *ProjectResponse {
	return &ProjectResponse{
		ID:        project.ID,
		CreatedAt: project.CreatedAt.Format(time.RFC3339),
		UpdatedAt: project.UpdatedAt.Format(time.RFC3339),
		Slug:      project.Slug,
		Name:      project.Name,
	}
}

type ListProjectsResponseProjectLocation struct {
	Name *string  `json:"name" example:"Office"`
	Lat  *float64 `json:"lat" example:"52.520008"`
	Lon  *float64 `json:"lon" example:"13.404954"`
}

func ToListProjectsResponseProjectLocation(project *database.ListProjectsWithLocationsRow) *ListProjectsResponseProjectLocation {
	if project.LocationName == nil {
		return nil
	}

	return &ListProjectsResponseProjectLocation{
		Name: project.LocationName,
		Lat:  project.LocationLat,
		Lon:  project.LocationLon,
	}
}

type listProjectsResponseProject struct {
	ID        int64                                `json:"id" example:"1"`
	CreatedAt string                               `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string                               `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Slug      string                               `json:"slug" example:"project-x"`
	Name      string                               `json:"name" example:"Project X"`
	Location  *ListProjectsResponseProjectLocation `json:"location"`
}

func toListProjectsResponseProject(project *database.ListProjectsWithLocationsRow) *listProjectsResponseProject {
	return &listProjectsResponseProject{
		ID:        project.ID,
		CreatedAt: project.CreatedAt.Format(time.RFC3339),
		UpdatedAt: project.UpdatedAt.Format(time.RFC3339),
		Slug:      project.Slug,
		Name:      project.Name,
		Location:  ToListProjectsResponseProjectLocation(project),
	}
}

type ListProjectsResponse struct {
	Projects []listProjectsResponseProject `json:"projects"`
	Total    int64                         `json:"total" example:"1"`
	Limit    int64                         `json:"limit" example:"100"`
	Offset   int64                         `json:"offset" example:"0"`
}

func (l *ListProjectsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToListProjectsResponse(result *FindAllProjectsResult) *ListProjectsResponse {
	listProjectsResponseProjects := make([]listProjectsResponseProject, len(result.Projects))
	for i, project := range result.Projects {
		listProjectsResponseProjects[i] = *toListProjectsResponseProject(project)
	}
	return &ListProjectsResponse{
		Projects: listProjectsResponseProjects,
		Total:    result.Total,
		Limit:    result.Limit,
		Offset:   result.Offset,
	}
}
