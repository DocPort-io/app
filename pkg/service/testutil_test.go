package service

import (
	"app/pkg/database"
	"app/pkg/storage"
	"context"
	"database/sql"
	"errors"
	"io"
	"log"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

// SpyFileStorage is a minimal spy that records method calls.
type SpyFileStorage struct {
	Calls []string
}

func NewSpyFileStorage() *SpyFileStorage { return &SpyFileStorage{} }

func (s *SpyFileStorage) Save(ctx context.Context, relativePath string, data io.Reader) error {
	s.Calls = append(s.Calls, "Save")
	return nil
}

func (s *SpyFileStorage) Retrieve(ctx context.Context, relativePath string) (io.ReadCloser, error) {
	s.Calls = append(s.Calls, "Retrieve")
	return nil, nil
}

func (s *SpyFileStorage) Delete(ctx context.Context, relativePath string) error {
	s.Calls = append(s.Calls, "Delete")
	return nil
}

func (s *SpyFileStorage) List(ctx context.Context, root string) ([]storage.ObjectInfo, error) {
	s.Calls = append(s.Calls, "List")
	return nil, nil
}

func (s *SpyFileStorage) Walk(ctx context.Context, root string, walkFunc storage.WalkFunc) error {
	s.Calls = append(s.Calls, "Walk")
	return nil
}

// setupTestDb prepares an in-memory SQLite database with migrations applied.
func setupTestDb(t *testing.T) (*sql.DB, *database.Queries) {
	t.Helper()
	db, err := sql.Open("sqlite", "file::memory:?_pragma=foreign_keys(ON)")
	if err != nil {
		log.Fatalf("failed to open database: %s\n", err)
	}
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatalf("failed to create database driver: %s\n", err)
	}
	migrations, err := migrate.NewWithDatabaseInstance("file://../../migrations", "sqlite", driver)
	if err != nil {
		log.Fatalf("failed to create migration source: %s\n", err)
	}
	if err := migrations.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("failed to run migrations: %s\n", err)
	}
	queries := database.New(db)
	return db, queries
}

// setupFileService wires a FileService using in-memory DB and a spy storage.
func setupFileService(t *testing.T) (*FileService, *sql.DB, *database.Queries, *SpyFileStorage) {
	t.Helper()
	db, queries := setupTestDb(t)
	spy := NewSpyFileStorage()
	fileService := NewFileService(queries, spy)
	return fileService, db, queries, spy
}

// setupVersionService wires a VersionService using in-memory DB and a spy storage.
func setupVersionService(t *testing.T) (*VersionService, *FileService, *database.Queries, *sql.DB, *SpyFileStorage) {
	t.Helper()
	fileSvc, db, queries, spy := setupFileService(t)
	vs := NewVersionService(queries)
	return vs, fileSvc, queries, db, spy
}

// setupProjectService wires a ProjectService on an in-memory test DB.
func setupProjectService(t *testing.T) (*ProjectService, *sql.DB) {
	t.Helper()
	db, queries := setupTestDb(t)
	svc := NewProjectService(queries)
	return svc, db
}

// seedProject creates a project row required for versions FKs.
func seedProject(t *testing.T, q *database.Queries, slug, name string) *database.Project {
	t.Helper()
	p, err := q.CreateProject(context.Background(), &database.CreateProjectParams{Slug: slug, Name: name})
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	return p
}

// seedVersion creates a version under a project.
func seedVersion(t *testing.T, q *database.Queries, projectID int64, name string, desc *string) *database.Version {
	t.Helper()
	v, err := q.CreateVersion(context.Background(), &database.CreateVersionParams{Name: name, Description: desc, ProjectID: projectID})
	if err != nil {
		t.Fatalf("failed to create version: %v", err)
	}
	return v
}
