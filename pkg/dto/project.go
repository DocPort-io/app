package dto

import (
	"app/pkg/model"
	"time"
)

type CreateProjectDto struct {
	Slug string `json:"slug" binding:"required" example:"project-x"`
	Name string `json:"name" binding:"required" example:"Project X"`
}

func (dto CreateProjectDto) ToModel() *model.Project {
	return &model.Project{
		Slug: dto.Slug,
		Name: dto.Name,
	}
}

type ProjectResponseDto struct {
	ID        uint   `json:"id" example:"1"`
	CreatedAt string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Slug      string `json:"slug" example:"project-x"`
	Name      string `json:"name" example:"Project X"`
}

func ToProjectResponse(project *model.Project) *ProjectResponseDto {
	return &ProjectResponseDto{
		ID:        project.ID,
		CreatedAt: project.CreatedAt.Format(time.RFC3339),
		UpdatedAt: project.UpdatedAt.Format(time.RFC3339),
		Slug:      project.Slug,
		Name:      project.Name,
	}
}

type ListProjectsResponseDto struct {
	Projects []ProjectResponseDto `json:"projects"`
	Total    int64                `json:"total" example:"1"`
}

func ToListProjectsResponse(projects []model.Project, total int64) *ListProjectsResponseDto {
	projectResponseDtos := make([]ProjectResponseDto, len(projects))
	for i, project := range projects {
		projectResponseDtos[i] = *ToProjectResponse(&project)
	}
	return &ListProjectsResponseDto{
		Projects: projectResponseDtos,
		Total:    total,
	}
}
