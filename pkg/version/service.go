package version

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"context"
	"database/sql"
	"errors"
)

type VersionService interface {
	FindAllVersions(ctx context.Context, params *FindAllVersionsParams) (*FindAllVersionsResult, error)
	FindVersionById(ctx context.Context, params *FindVersionByIdParams) (*FindVersionByIdResult, error)
	CreateVersion(ctx context.Context, params *CreateVersionParams) (*CreateVersionResult, error)
	UpdateVersion(ctx context.Context, params *UpdateVersionParams) (*UpdateVersionResult, error)
	DeleteVersion(ctx context.Context, params *DeleteVersionParams) error
	AttachFileToVersion(ctx context.Context, params *AttachFileToVersionParams) error
	DetachFileFromVersion(ctx context.Context, params *DetachFileFromVersionParams) error
}

type versionServiceImpl struct {
	queries *database.Queries
}

func NewVersionService(queries *database.Queries) VersionService {
	return &versionServiceImpl{queries: queries}
}

func (s *versionServiceImpl) FindAllVersions(ctx context.Context, params *FindAllVersionsParams) (*FindAllVersionsResult, error) {
	versions, err := s.queries.ListVersionsByProjectId(ctx, &database.ListVersionsByProjectIdParams{
		ProjectID: params.ProjectID,
		Limit:     params.Limit,
		Offset:    params.Offset,
	})
	if err != nil {
		return nil, err
	}

	count, err := s.queries.CountVersionsByProjectId(ctx, params.ProjectID)
	if err != nil {
		return nil, err
	}

	return &FindAllVersionsResult{Versions: versions, Total: count, Limit: params.Limit, Offset: params.Offset}, nil
}

func (s *versionServiceImpl) FindVersionById(ctx context.Context, params *FindVersionByIdParams) (*FindVersionByIdResult, error) {
	version, err := s.queries.GetVersion(ctx, params.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	return &FindVersionByIdResult{Version: version}, nil
}

func (s *versionServiceImpl) CreateVersion(ctx context.Context, params *CreateVersionParams) (*CreateVersionResult, error) {
	version, err := s.queries.CreateVersion(ctx, &database.CreateVersionParams{
		Name:        params.Name,
		Description: params.Description,
		ProjectID:   params.ProjectID,
	})
	if err != nil {
		return nil, err
	}

	return &CreateVersionResult{Version: version}, nil
}

func (s *versionServiceImpl) UpdateVersion(ctx context.Context, params *UpdateVersionParams) (*UpdateVersionResult, error) {
	version, err := s.queries.UpdateVersion(ctx, &database.UpdateVersionParams{
		Name:        params.Name,
		Description: params.Description,
		ID:          params.ID,
	})
	if err != nil {
		return nil, err
	}

	return &UpdateVersionResult{Version: version}, nil
}

func (s *versionServiceImpl) DeleteVersion(ctx context.Context, params *DeleteVersionParams) error {
	err := s.queries.DeleteVersion(ctx, params.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *versionServiceImpl) AttachFileToVersion(ctx context.Context, params *AttachFileToVersionParams) error {
	err := s.queries.AttachFileToVersion(ctx, &database.AttachFileToVersionParams{
		VersionID: params.VersionID,
		FileID:    params.FileID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *versionServiceImpl) DetachFileFromVersion(ctx context.Context, params *DetachFileFromVersionParams) error {
	err := s.queries.DetachFileFromVersion(ctx, &database.DetachFileFromVersionParams{
		VersionID: params.VersionID,
		FileID:    params.FileID,
	})
	if err != nil {
		return err
	}

	return nil
}
