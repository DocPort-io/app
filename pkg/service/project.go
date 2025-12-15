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
	FindAllProjects(ctx context.Context, params *dto.FindAllProjectsParams) (*dto.FindAllProjectsResult, error)
	FindProjectById(ctx context.Context, params *dto.FindProjectByIdParams) (*dto.FindProjectByIdResult, error)
	CreateProject(ctx context.Context, params *dto.CreateProjectParams) (*dto.CreateProjectResult, error)
	UpdateProject(ctx context.Context, params *dto.UpdateProjectParams) (*dto.UpdateProjectResult, error)
	DeleteProject(ctx context.Context, params *dto.DeleteProjectParams) error
}

type projectServiceImpl struct {
	queries *database.Queries
}

func NewProjectService(queries *database.Queries) ProjectService {
	return &projectServiceImpl{queries: queries}
}

func (s *projectServiceImpl) FindAllProjects(ctx context.Context, params *dto.FindAllProjectsParams) (*dto.FindAllProjectsResult, error) {
	projects, err := s.queries.ListProjectsWithLocations(ctx, &database.ListProjectsWithLocationsParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, err
	}

	count, err := s.queries.CountProjects(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.FindAllProjectsResult{
		Projects: projects,
		Total:    count,
		Limit:    params.Limit,
		Offset:   params.Offset,
	}, nil
}

func (s *projectServiceImpl) FindProjectById(ctx context.Context, params *dto.FindProjectByIdParams) (*dto.FindProjectByIdResult, error) {
	project, err := s.queries.GetProject(ctx, params.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	return &dto.FindProjectByIdResult{Project: project}, nil
}

func (s *projectServiceImpl) CreateProject(ctx context.Context, params *dto.CreateProjectParams) (*dto.CreateProjectResult, error) {
	project, err := s.queries.CreateProject(ctx, &database.CreateProjectParams{
		Slug:       params.Slug,
		Name:       params.Name,
		LocationID: nil,
	})
	if err != nil {
		return nil, err
	}

	return &dto.CreateProjectResult{Project: project}, nil
}

func (s *projectServiceImpl) UpdateProject(ctx context.Context, params *dto.UpdateProjectParams) (*dto.UpdateProjectResult, error) {
	project, err := s.queries.UpdateProject(ctx, &database.UpdateProjectParams{
		Slug:       params.Slug,
		Name:       params.Name,
		LocationID: nil,
		ID:         params.ID,
	})
	if err != nil {
		return nil, err
	}

	return &dto.UpdateProjectResult{Project: project}, nil
}

func (s *projectServiceImpl) DeleteProject(ctx context.Context, params *dto.DeleteProjectParams) error {
	err := s.queries.DeleteProject(ctx, params.ID)
	if err != nil {
		return err
	}

	return nil
}
