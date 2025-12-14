package service

import (
	"app/pkg/database"
	"app/pkg/dto"
	"context"
)

type VersionService interface {
	FindAllVersions(ctx context.Context, projectId int64) ([]*database.Version, int64, error)
	FindVersionById(ctx context.Context, id int64) (*database.Version, error)
	CreateVersion(ctx context.Context, dto *dto.CreateVersionDto) (*database.Version, error)
	UpdateVersion(ctx context.Context, id int64, dto *dto.UpdateVersionDto) (*database.Version, error)
	DeleteVersion(ctx context.Context, id int64) error
	AttachFileToVersion(ctx context.Context, id int64, attachFileToVersionDto *dto.AttachFileToVersionDto) error
	DetachFileFromVersion(ctx context.Context, id int64, detachFileFromVersionDto *dto.DetachFileFromVersionDto) error
}

type versionServiceImpl struct {
	queries *database.Queries
}

func NewVersionService(queries *database.Queries) VersionService {
	return &versionServiceImpl{queries: queries}
}

func (s *versionServiceImpl) FindAllVersions(ctx context.Context, projectId int64) ([]*database.Version, int64, error) {
	versions, err := s.queries.ListVersionsByProjectId(ctx, projectId)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.queries.CountVersionsByProjectId(ctx, projectId)
	if err != nil {
		return nil, 0, err
	}

	return versions, count, nil
}

func (s *versionServiceImpl) FindVersionById(ctx context.Context, id int64) (*database.Version, error) {
	version, err := s.queries.GetVersion(ctx, id)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (s *versionServiceImpl) CreateVersion(ctx context.Context, dto *dto.CreateVersionDto) (*database.Version, error) {
	version, err := s.queries.CreateVersion(ctx, &database.CreateVersionParams{
		Name:        dto.Name,
		Description: dto.Description,
		ProjectID:   dto.ProjectId,
	})
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (s *versionServiceImpl) UpdateVersion(ctx context.Context, id int64, dto *dto.UpdateVersionDto) (*database.Version, error) {
	version, err := s.queries.UpdateVersion(ctx, &database.UpdateVersionParams{
		Name:        dto.Name,
		Description: dto.Description,
		ID:          id,
	})
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (s *versionServiceImpl) DeleteVersion(ctx context.Context, id int64) error {
	err := s.queries.DeleteVersion(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *versionServiceImpl) AttachFileToVersion(ctx context.Context, id int64, attachFileToVersionDto *dto.AttachFileToVersionDto) error {
	err := s.queries.AttachFileToVersion(ctx, &database.AttachFileToVersionParams{
		VersionID: id,
		FileID:    attachFileToVersionDto.FileId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *versionServiceImpl) DetachFileFromVersion(ctx context.Context, id int64, detachFileFromVersionDto *dto.DetachFileFromVersionDto) error {
	err := s.queries.DetachFileFromVersion(ctx, &database.DetachFileFromVersionParams{
		VersionID: id,
		FileID:    detachFileFromVersionDto.FileId,
	})
	if err != nil {
		return err
	}

	return nil
}
