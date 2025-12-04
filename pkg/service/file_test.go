package service

import (
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/storage"
	"context"
	"database/sql"
	"errors"
	"io"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

type SpyFileStorage struct {
	Calls []string
}

func NewSpyFileStorage() *SpyFileStorage {
	return &SpyFileStorage{}
}

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

func createTempFile(t *testing.T) string {
	t.Helper()

	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}

	bytesWritten, err := file.Write([]byte("Hello, world!"))
	if err != nil {
		t.Fatal(err)
	}

	if bytesWritten != len("Hello, world!") {
		t.Fatalf("failed to write all bytes")
	}

	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.Remove(file.Name())
	})

	return file.Name()
}

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

	err = migrations.Up()
	if err != nil {
		log.Fatalf("failed to run migrations: %s\n", err)
	}

	queries := database.New(db)

	return db, queries
}

func setupFileService(t *testing.T) (*FileService, *sql.DB, *database.Queries, *SpyFileStorage) {
	t.Helper()

	db, queries := setupTestDb(t)
	fileStorage := NewSpyFileStorage()
	fileService := NewFileService(queries, fileStorage)
	return fileService, db, queries, fileStorage
}

func TestFileService_FindAllFiles(t *testing.T) {
	// Arrange
	service, _, queries, _ := setupFileService(t)

	var testFile = &database.CreateFileParams{
		Name: "test-file.pdf",
		Size: 1024,
		Path: "files/abcd.pdf",
	}

	_, err := queries.CreateFile(t.Context(), testFile)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Act
	result, err := service.FindAllFiles(t.Context(), nil)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Errorf("expected at least one file, got none")
	}
}

func TestFileService_CreateFile(t *testing.T) {
	// Arrange
	service, _, _, fileStorage := setupFileService(t)
	tempFile := createTempFile(t)

	// Act
	file := &dto.CreateFileDto{
		Name: "test-file.pdf",
		Size: 1024,
		Path: tempFile,
	}

	result, err := service.CreateFile(t.Context(), *file)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Errorf("expected a file, got nil")
	}

	if len(fileStorage.Calls) == 0 {
		t.Errorf("expected at least one call to fileStorage.Save, got none")
	} else if fileStorage.Calls[0] != "Save" {
		t.Errorf("expected call to fileStorage.Save, got %s", fileStorage.Calls[0])
	}
}

func TestFileService_FindFileById(t *testing.T) {
	// Arrange
	service, _, queries, _ := setupFileService(t)

	testFile := &database.CreateFileParams{
		Name: "find-me.pdf",
		Size: 2048,
		Path: "files/find-me.pdf",
	}

	file, err := queries.CreateFile(t.Context(), testFile)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Act
	got, err := service.FindFileById(t.Context(), &file.ID)

	// Assert
	if err != nil {
		t.Fatalf("FindFileById returned error: %v", err)
	}
	if got == nil {
		t.Fatalf("expected a file, got nil")
	}
	if got.ID != file.ID || got.Name != testFile.Name || got.Size != testFile.Size {
		t.Errorf("returned file does not match: got %+v, want %+v", got, testFile)
	}
}

func TestFileService_FindFileById_NotFound(t *testing.T) {
	// Arrange
	service, _, _, _ := setupFileService(t)

	// Act
	var fileId int64 = 999999
	_, err := service.FindFileById(t.Context(), &fileId)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected ErrRecordNotFound, got %v", err)
	}
}
