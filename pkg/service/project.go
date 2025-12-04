package service

import (
	"app/pkg/database"
	"app/pkg/dto"
	"context"
)

type ProjectService struct {
	queries *database.Queries
}

func NewProjectService(queries *database.Queries) *ProjectService {
	return &ProjectService{queries: queries}
}

func (s *ProjectService) FindAllProjects(ctx context.Context) ([]*database.ListProjectsWithLocationsRow, error) {
	projects, err := s.queries.ListProjectsWithLocations(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *ProjectService) FindProjectById(ctx context.Context, id *int64) (*database.Project, error) {
	project, err := s.queries.GetProject(ctx, *id)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) CreateProject(ctx context.Context, dto dto.CreateProjectDto) (*database.Project, error) {
	project, err := s.queries.CreateProject(ctx, &database.CreateProjectParams{
		Slug:       dto.Slug,
		Name:       dto.Name,
		LocationID: nil,
	})
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) UpdateProject(ctx context.Context, id *int64, dto dto.UpdateProjectDto) (*database.Project, error) {
	project, err := s.queries.UpdateProject(ctx, &database.UpdateProjectParams{
		Slug:       dto.Slug,
		Name:       dto.Name,
		LocationID: nil,
		ID:         *id,
	})
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, id *int64) error {
	err := s.queries.DeleteProject(ctx, *id)
	if err != nil {
		return err
	}

	return nil
}
