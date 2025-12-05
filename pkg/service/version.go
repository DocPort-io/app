package service

import (
	"app/pkg/database"
	"app/pkg/dto"
	"context"
)

type VersionService struct {
	queries *database.Queries
}

func NewVersionService(queries *database.Queries) *VersionService {
	return &VersionService{queries: queries}
}

func (s *VersionService) FindAllVersions(ctx context.Context, projectId *int64) ([]*database.Version, error) {
	if projectId != nil {
		versions, err := s.queries.ListVersionsByProjectId(ctx, *projectId)
		if err != nil {
			return nil, err
		}
		return versions, nil
	}

	versions, err := s.queries.ListVersions(ctx)
	if err != nil {
		return nil, err
	}
	return versions, nil
}

func (s *VersionService) FindVersionById(ctx context.Context, id *int64) (*database.Version, error) {
	version, err := s.queries.GetVersion(ctx, *id)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (s *VersionService) CreateVersion(ctx context.Context, dto dto.CreateVersionDto) (*database.Version, error) {
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

func (s *VersionService) UpdateVersion(ctx context.Context, id *int64, dto dto.UpdateVersionDto) (*database.Version, error) {
	version, err := s.queries.UpdateVersion(ctx, &database.UpdateVersionParams{
		Name:        dto.Name,
		Description: dto.Description,
		ID:          *id,
	})
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (s *VersionService) DeleteVersion(ctx context.Context, id *int64) error {
	err := s.queries.DeleteVersion(ctx, *id)
	if err != nil {
		return err
	}

	return nil
}

func (s *VersionService) AttachFileToVersion(ctx context.Context, id *int64, attachFileToVersionDto dto.AttachFileToVersionDto) error {
	err := s.queries.AttachFileToVersion(ctx, &database.AttachFileToVersionParams{
		VersionID: *id,
		FileID:    attachFileToVersionDto.FileId,
	})
	if err != nil {
		return err
	}

	return nil
}
