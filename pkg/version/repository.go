package version

import (
	"app/pkg/database"
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

var (
	ErrVersionNotFound            = errors.New("version not found")
	ErrVersionAlreadyExists       = errors.New("version already exists")
	ErrVersionFileAlreadyAttached = errors.New("version file already attached")
)

type Repository interface {
	GetById(ctx context.Context, id int64) (Version, error)
	List(ctx context.Context, projectId *int64, limit, offset int64) ([]Version, error)
	Create(ctx context.Context, version Version) (Version, error)
	Update(ctx context.Context, version Version) (Version, error)
	Delete(ctx context.Context, id int64) error
	AttachFile(ctx context.Context, id int64, fileId int64) error
	DetachFile(ctx context.Context, id int64, fileId int64) error
}

type repository struct {
	queries *database.Queries
}

func NewRepository(queries *database.Queries) Repository {
	return &repository{queries: queries}
}

func (r *repository) GetById(ctx context.Context, id int64) (Version, error) {
	row, err := r.queries.GetVersion(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return Version{}, ErrVersionNotFound
	}
	if err != nil {
		return Version{}, err
	}
	return toVersion(row), nil
}

func (r *repository) List(ctx context.Context, projectId *int64, limit, offset int64) ([]Version, error) {
	rows, err := r.queries.ListVersions(ctx, &database.ListVersionsParams{
		ProjectId: projectId,
		Offset:    offset,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}
	versions := make([]Version, len(rows))
	for i, row := range rows {
		versions[i] = toVersion(row)
	}
	return versions, nil
}

func (r *repository) Create(ctx context.Context, version Version) (Version, error) {
	row, err := r.queries.CreateVersion(ctx, &database.CreateVersionParams{
		Name:        version.Name,
		Description: version.Description,
		ProjectID:   version.ProjectID,
	})
	if err != nil {
		if isPgUniqueViolation(err) {
			return Version{}, ErrVersionAlreadyExists
		}
		return Version{}, err
	}
	return toVersion(row), nil
}

func (r *repository) Update(ctx context.Context, version Version) (Version, error) {
	row, err := r.queries.UpdateVersion(ctx, &database.UpdateVersionParams{
		ID:          version.ID,
		Name:        version.Name,
		Description: version.Description,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Version{}, ErrVersionNotFound
	}
	if err != nil {
		return Version{}, err
	}
	return toVersion(row), nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	err := r.queries.DeleteVersion(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrVersionNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) AttachFile(ctx context.Context, id int64, fileId int64) error {
	err := r.queries.AttachFileToVersion(ctx, &database.AttachFileToVersionParams{
		VersionID: id,
		FileID:    fileId,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrVersionNotFound
	}
	if err != nil {
		if isPgUniqueViolation(err) {
			return ErrVersionFileAlreadyAttached
		}
		return err
	}
	return nil
}

func (r *repository) DetachFile(ctx context.Context, id int64, fileId int64) error {
	err := r.queries.DetachFileFromVersion(ctx, &database.DetachFileFromVersionParams{
		VersionID: id,
		FileID:    fileId,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrVersionNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func toVersion(row *database.Version) Version {
	return Version{
		ID:          row.ID,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
		Name:        row.Name,
		Description: row.Description,
		ProjectID:   row.ProjectID,
	}
}

func isPgUniqueViolation(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "23505") || strings.Contains(err.Error(), "unique"))
}
