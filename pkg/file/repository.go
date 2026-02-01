package file

import (
	"app/pkg/database"
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

var (
	ErrFileNotFound     = errors.New("file not found")
	ErrFileAlreadyExist = errors.New("file already exist")
)

type Repository interface {
	GetById(ctx context.Context, id int64) (File, error)
	List(ctx context.Context, versionId *int64, limit, offset int64) ([]File, error)
	Create(ctx context.Context, file File) (File, error)
	Update(ctx context.Context, file File) (File, error)
	Delete(ctx context.Context, id int64) error
}

type repository struct {
	queries *database.Queries
}

func NewRepository(queries *database.Queries) Repository {
	return &repository{queries: queries}
}

func (r *repository) GetById(ctx context.Context, id int64) (File, error) {
	row, err := r.queries.GetFile(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return File{}, ErrFileNotFound
	}
	if err != nil {
		return File{}, err
	}
	return toFile(row), nil
}

func (r *repository) List(ctx context.Context, versionId *int64, limit, offset int64) ([]File, error) {
	rows, err := r.queries.ListFiles(ctx, &database.ListFilesParams{
		VersionId: versionId,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		return nil, err
	}
	files := make([]File, len(rows))
	for i, row := range rows {
		files[i] = toFile(row)
	}
	return files, nil
}

func (r *repository) Create(ctx context.Context, file File) (File, error) {
	row, err := r.queries.CreateFile(ctx, &database.CreateFileParams{
		Name:       file.Name,
		Size:       file.Size,
		Path:       file.Path,
		MimeType:   file.MimeType,
		IsComplete: file.IsComplete,
	})
	if err != nil {
		if isPgUniqueViolation(err) {
			return File{}, ErrFileAlreadyExist
		}
		return File{}, err
	}
	return toFile(row), nil
}

func (r *repository) Update(ctx context.Context, file File) (File, error) {
	row, err := r.queries.UpdateFile(ctx, &database.UpdateFileParams{
		ID:         file.ID,
		Name:       file.Name,
		Size:       file.Size,
		Path:       file.Path,
		MimeType:   file.MimeType,
		IsComplete: file.IsComplete,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return File{}, ErrFileNotFound
	}
	if err != nil {
		return File{}, err
	}
	return toFile(row), nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	err := r.queries.DeleteFile(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrFileNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func toFile(row *database.File) File {
	return File{
		ID:         row.ID,
		CreatedAt:  row.CreatedAt.Time,
		UpdatedAt:  row.UpdatedAt.Time,
		Name:       row.Name,
		Size:       row.Size,
		Path:       row.Path,
		MimeType:   row.MimeType,
		IsComplete: row.IsComplete,
	}
}

func isPgUniqueViolation(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "23505") || strings.Contains(err.Error(), "unique"))
}
