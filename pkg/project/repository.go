package project

import (
	"app/pkg/database"
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

var (
	ErrProjectNotFound      = errors.New("project not found")
	ErrProjectAlreadyExists = errors.New("project already exists")
)

type Repository interface {
	GetById(ctx context.Context, id int64) (Project, error)
	List(ctx context.Context, limit, offset int64) ([]Project, error)
	Create(ctx context.Context, project Project) (Project, error)
	Update(ctx context.Context, project Project) (Project, error)
	Delete(ctx context.Context, id int64) error
}

type repository struct {
	queries *database.Queries // db or any other dependencies
}

func NewRepository(queries *database.Queries) Repository {
	return &repository{queries: queries}
}

func (r *repository) GetById(ctx context.Context, id int64) (Project, error) {
	row, err := r.queries.GetProject(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return Project{}, ErrProjectNotFound
	}
	if err != nil {
		return Project{}, err
	}
	return toProject(row), nil
}

func (r *repository) List(ctx context.Context, limit, offset int64) ([]Project, error) {
	rows, err := r.queries.ListProjects(ctx, &database.ListProjectsParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	projects := make([]Project, len(rows))
	for i, row := range rows {
		projects[i] = toProject(row)
	}
	return projects, nil
}

func (r *repository) Create(ctx context.Context, project Project) (Project, error) {
	row, err := r.queries.CreateProject(ctx, &database.CreateProjectParams{
		Slug: project.Slug,
		Name: project.Name,
	})
	if err != nil {
		if isPgUniqueViolation(err) {
			return Project{}, ErrProjectAlreadyExists
		}
		return Project{}, err
	}
	return toProject(row), nil
}

func (r *repository) Update(ctx context.Context, project Project) (Project, error) {
	row, err := r.queries.UpdateProject(ctx, &database.UpdateProjectParams{
		ID:   project.ID,
		Slug: project.Slug,
		Name: project.Name,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Project{}, ErrProjectNotFound
	}
	if err != nil {
		return Project{}, err
	}
	return toProject(row), nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	err := r.queries.DeleteProject(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrProjectNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func toProject(row *database.Project) Project {
	return Project{
		ID:        row.ID,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		Slug:      row.Slug,
		Name:      row.Name,
	}
}
func isPgUniqueViolation(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "23505") || strings.Contains(err.Error(), "unique"))
}
