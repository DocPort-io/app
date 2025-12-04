package dto

import (
	"app/pkg/database"
	"time"
)

type CreateProjectDto struct {
	Slug string `json:"slug" binding:"required" example:"project-x"`
	Name string `json:"name" binding:"required" example:"Project X"`
}

type UpdateProjectDto struct {
	Slug string `json:"slug" example:"project-x"`
	Name string `json:"name" example:"Project X"`
}

type ProjectResponseDto struct {
	ID        int64  `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Slug      string `json:"slug" example:"project-x"`
	Name      string `json:"name" example:"Project X"`
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
}

func ToListProjectsResponse(projects []*database.ListProjectsWithLocationsRow, total int64) *ListProjectsResponseDto {
	projectResponseProjectDtos := make([]ListProjectsResponseProjectDto, len(projects))
	for i, project := range projects {
		projectResponseProjectDtos[i] = *ToListProjectsResponseProject(project)
	}
	return &ListProjectsResponseDto{
		Projects: projectResponseProjectDtos,
		Total:    total,
	}
}
