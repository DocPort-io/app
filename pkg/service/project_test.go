package service

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/paginate"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test: FindAllProjects paginates.
func TestProjectService_FindAllProjects_Paginates(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewProjectService(queries)
	_, _ = queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p-1", Name: "Project 1"})
	_, _ = queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p-2", Name: "Project 2"})

	res, err := svc.FindAllProjects(context.Background(), &dto.FindAllProjectsParams{Pagination: &paginate.Pagination{Limit: 1, Offset: 0}})
	assert.NoError(t, err)
	assert.Equal(t, int64(2), res.Total)
	assert.Len(t, res.Projects, 1)
}

// Test: FindProjectById returns project by ID.
func TestProjectService_FindProjectById_Finds(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewProjectService(queries)
	p, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p-1", Name: "Project 1"})

	got, err := svc.FindProjectById(context.Background(), &dto.FindProjectByIdParams{ID: p.ID})
	assert.NoError(t, err)
	assert.Equal(t, p.ID, got.Project.ID)
}

// Test: FindProjectById returns not found error when missing.
func TestProjectService_FindProjectById_NotFound_ReturnsErr(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewProjectService(queries)

	_, err := svc.FindProjectById(context.Background(), &dto.FindProjectByIdParams{ID: 999})
	assert.Error(t, err)
	assert.ErrorIs(t, err, apperrors.ErrNotFound)
}

// Test: CreateProject creates a new project.
func TestProjectService_CreateProject_Creates(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewProjectService(queries)

	created, err := svc.CreateProject(context.Background(), &dto.CreateProjectParams{Slug: "p-3", Name: "Project 3"})
	assert.NoError(t, err)
	assert.NotZero(t, created.Project.ID)
}

// Test: UpdateProject updates fields.
func TestProjectService_UpdateProject_Updates(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewProjectService(queries)
	p, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p-1", Name: "Project 1"})

	upd, err := svc.UpdateProject(context.Background(), &dto.UpdateProjectParams{ID: p.ID, Slug: "p-1b", Name: "Project 1b"})
	assert.NoError(t, err)
	assert.Equal(t, "p-1b", upd.Project.Slug)
	assert.Equal(t, "Project 1b", upd.Project.Name)
}

// Test: DeleteProject deletes the project.
func TestProjectService_DeleteProject_Deletes(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewProjectService(queries)
	p, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p-2", Name: "Project 2"})

	err := svc.DeleteProject(context.Background(), &dto.DeleteProjectParams{ID: p.ID})
	assert.NoError(t, err)

	_, err = queries.GetProject(context.Background(), p.ID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}
