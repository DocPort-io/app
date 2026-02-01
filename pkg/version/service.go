package version

import (
	"context"
)

type Service interface {
	GetById(ctx context.Context, id int64) (Version, error)
	List(ctx context.Context, projectId *int64, limit, offset int64) ([]Version, error)
	Create(ctx context.Context, req CreateVersionRequest) (Version, error)
	Update(ctx context.Context, id int64, req UpdateVersionRequest) (Version, error)
	Delete(ctx context.Context, id int64) error
	AttachFile(ctx context.Context, id int64, req AttachFileRequest) error
	DetachFile(ctx context.Context, id int64, req DetachFileRequest) error
}

type service struct {
	repository Repository
}

func NewVersionService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetById(ctx context.Context, id int64) (Version, error) {
	return s.repository.GetById(ctx, id)
}

func (s *service) List(ctx context.Context, projectId *int64, limit, offset int64) ([]Version, error) {
	return s.repository.List(ctx, projectId, limit, offset)
}

func (s *service) Create(ctx context.Context, req CreateVersionRequest) (Version, error) {
	version := Version{
		Name:        req.Name,
		Description: req.Description,
		ProjectID:   req.ProjectId,
	}
	return s.repository.Create(ctx, version)
}

func (s *service) Update(ctx context.Context, id int64, req UpdateVersionRequest) (Version, error) {
	version := Version{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}
	return s.repository.Update(ctx, version)
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) AttachFile(ctx context.Context, id int64, req AttachFileRequest) error {
	return s.repository.AttachFile(ctx, id, req.FileID)
}

func (s *service) DetachFile(ctx context.Context, id int64, req DetachFileRequest) error {
	return s.repository.DetachFile(ctx, id, req.FileID)
}
