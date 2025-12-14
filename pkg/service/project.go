package service

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"context"
	"database/sql"
	"errors"
)

type ProjectService interface {
	FindAllProjects(ctx context.Context) ([]*database.ListProjectsWithLocationsRow, int64, error)
	FindProjectById(ctx context.Context, id int64) (*database.Project, error)
	CreateProject(ctx context.Context, dto *dto.CreateProjectDto) (*database.Project, error)
	UpdateProject(ctx context.Context, id int64, dto *dto.UpdateProjectDto) (*database.Project, error)
	DeleteProject(ctx context.Context, id int64) error
}

type projectServiceImpl struct {
	queries *database.Queries
}

func NewProjectService(queries *database.Queries) ProjectService {
	return &projectServiceImpl{queries: queries}
}

func (s *projectServiceImpl) FindAllProjects(ctx context.Context) ([]*database.ListProjectsWithLocationsRow, int64, error) {
	projects, err := s.queries.ListProjectsWithLocations(ctx)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.queries.CountProjects(ctx)
	if err != nil {
		return nil, 0, err
	}

	return projects, count, nil
}

func (s *projectServiceImpl) FindProjectById(ctx context.Context, id int64) (*database.Project, error) {
	project, err := s.queries.GetProject(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	return project, nil
}

func (s *projectServiceImpl) CreateProject(ctx context.Context, dto *dto.CreateProjectDto) (*database.Project, error) {
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

func (s *projectServiceImpl) UpdateProject(ctx context.Context, id int64, dto *dto.UpdateProjectDto) (*database.Project, error) {
	project, err := s.queries.UpdateProject(ctx, &database.UpdateProjectParams{
		Slug:       dto.Slug,
		Name:       dto.Name,
		LocationID: nil,
		ID:         id,
	})
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *projectServiceImpl) DeleteProject(ctx context.Context, id int64) error {
	err := s.queries.DeleteProject(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
