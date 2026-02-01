package project

import (
	"context"
)

type Service interface {
	GetById(ctx context.Context, id int64) (Project, error)
	List(ctx context.Context, limit, offset int64) ([]Project, error)
	CreateProject(ctx context.Context, req CreateProjectRequest) (Project, error)
	UpdateProject(ctx context.Context, id int64, req UpdateProjectRequest) (Project, error)
	DeleteProject(ctx context.Context, id int64) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetById(ctx context.Context, id int64) (Project, error) {
	return s.repository.GetById(ctx, id)
}

func (s *service) List(ctx context.Context, limit, offset int64) ([]Project, error) {
	return s.repository.List(ctx, limit, offset)
}

func (s *service) CreateProject(ctx context.Context, req CreateProjectRequest) (Project, error) {
	project := Project{
		Slug: req.Slug,
		Name: req.Name,
	}
	return s.repository.Create(ctx, project)
}

func (s *service) UpdateProject(ctx context.Context, id int64, req UpdateProjectRequest) (Project, error) {
	project := Project{
		ID:   id,
		Slug: req.Slug,
		Name: req.Name,
	}
	return s.repository.Update(ctx, project)
}

func (s *service) DeleteProject(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
