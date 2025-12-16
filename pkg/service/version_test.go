package service

import (
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/paginate"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test: CreateVersion creates a new version under project.
func TestVersionService_CreateVersion_Creates(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewVersionService(queries)
	proj, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p", Name: "Proj"})

	res, err := svc.CreateVersion(context.Background(), &dto.CreateVersionParams{Name: "v1", ProjectID: proj.ID})
	assert.NoError(t, err)
	assert.NotZero(t, res.Version.ID)
}

// Test: FindAllVersions paginates results for a project.
func TestVersionService_FindAllVersions_Paginates(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewVersionService(queries)
	proj, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p", Name: "Proj"})
	_, _ = svc.CreateVersion(context.Background(), &dto.CreateVersionParams{Name: "v1", ProjectID: proj.ID})
	_, _ = svc.CreateVersion(context.Background(), &dto.CreateVersionParams{Name: "v2", ProjectID: proj.ID})

	list, err := svc.FindAllVersions(context.Background(), &dto.FindAllVersionsParams{ProjectID: proj.ID, Pagination: &paginate.Pagination{Limit: 1, Offset: 0}})
	assert.NoError(t, err)
	assert.Equal(t, int64(2), list.Total)
	assert.Len(t, list.Versions, 1)
}

// Test: FindVersionById returns the version by id.
func TestVersionService_FindVersionById_Finds(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewVersionService(queries)
	proj, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p", Name: "Proj"})
	v, _ := svc.CreateVersion(context.Background(), &dto.CreateVersionParams{Name: "v1", ProjectID: proj.ID})

	got, err := svc.FindVersionById(context.Background(), &dto.FindVersionByIdParams{ID: v.Version.ID})
	assert.NoError(t, err)
	assert.Equal(t, v.Version.ID, got.Version.ID)
}

// Test: UpdateVersion updates fields.
func TestVersionService_UpdateVersion_Updates(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewVersionService(queries)
	proj, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p", Name: "Proj"})
	v, _ := svc.CreateVersion(context.Background(), &dto.CreateVersionParams{Name: "v1", ProjectID: proj.ID})

	desc := "desc"
	upd, err := svc.UpdateVersion(context.Background(), &dto.UpdateVersionParams{ID: v.Version.ID, Name: "v1.1", Description: &desc})
	assert.NoError(t, err)
	assert.Equal(t, "v1.1", upd.Version.Name)
}

// Test: AttachFileToVersion associates the file.
func TestVersionService_AttachFileToVersion_Attaches(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewVersionService(queries)
	proj, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p", Name: "Proj"})
	v, _ := svc.CreateVersion(context.Background(), &dto.CreateVersionParams{Name: "v1", ProjectID: proj.ID})
	f, _ := queries.CreateFile(context.Background(), "file.pdf")

	err := svc.AttachFileToVersion(context.Background(), &dto.AttachFileToVersionParams{VersionID: v.Version.ID, FileID: f.ID})
	assert.NoError(t, err)
}

// Test: DetachFileFromVersion removes the association.
func TestVersionService_DetachFileFromVersion_Detaches(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewVersionService(queries)
	proj, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p", Name: "Proj"})
	v, _ := svc.CreateVersion(context.Background(), &dto.CreateVersionParams{Name: "v1", ProjectID: proj.ID})
	f, _ := queries.CreateFile(context.Background(), "file.pdf")
	_ = svc.AttachFileToVersion(context.Background(), &dto.AttachFileToVersionParams{VersionID: v.Version.ID, FileID: f.ID})

	err := svc.DetachFileFromVersion(context.Background(), &dto.DetachFileFromVersionParams{VersionID: v.Version.ID, FileID: f.ID})
	assert.NoError(t, err)
}

// Test: DeleteVersion deletes the version.
func TestVersionService_DeleteVersion_Deletes(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewVersionService(queries)
	proj, _ := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p", Name: "Proj"})
	v, _ := svc.CreateVersion(context.Background(), &dto.CreateVersionParams{Name: "v1", ProjectID: proj.ID})

	err := svc.DeleteVersion(context.Background(), &dto.DeleteVersionParams{ID: v.Version.ID})
	assert.NoError(t, err)

	// Verify gone
	_, err = queries.GetVersion(context.Background(), v.Version.ID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}
