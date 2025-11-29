package service

import (
	"app/pkg/dto"
	"app/pkg/model"
	"app/pkg/storage"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
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

func setupTestDb(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(&model.File{})
	if err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}

	return db
}

func setupFileService(t *testing.T) (*FileService, *gorm.DB, *SpyFileStorage) {
	t.Helper()

	db := setupTestDb(t)
	fileStorage := NewSpyFileStorage()
	fileService := NewFileService(db, fileStorage)
	return fileService, db, fileStorage
}

func TestFileService_FindAllFiles(t *testing.T) {
	// Arrange
	service, db, _ := setupFileService(t)

	var testFile = &model.File{
		Name: "test-file.pdf",
		Size: 1024,
		Path: "files/abcd.pdf",
	}

	err := gorm.G[model.File](db).Create(t.Context(), testFile)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Act
	result, err := service.FindAllFiles(t.Context(), "")

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
	service, _, fileStorage := setupFileService(t)
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
	service, db, _ := setupFileService(t)

	testFile := &model.File{
		Name: "find-me.pdf",
		Size: 2048,
		Path: "files/find-me.pdf",
	}

	if err := gorm.G[model.File](db).Create(t.Context(), testFile); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Act
	idStr := fmt.Sprintf("%d", testFile.ID)
	got, err := service.FindFileById(t.Context(), idStr)

	// Assert
	if err != nil {
		t.Fatalf("FindFileById returned error: %v", err)
	}
	if got == nil {
		t.Fatalf("expected a file, got nil")
	}
	if got.ID != testFile.ID || got.Name != testFile.Name || got.Size != testFile.Size {
		t.Errorf("returned file does not match: got %+v, want %+v", got, testFile)
	}
}

func TestFileService_FindFileById_NotFound(t *testing.T) {
	// Arrange
	service, _, _ := setupFileService(t)

	// Act
	_, err := service.FindFileById(t.Context(), "999999")

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected ErrRecordNotFound, got %v", err)
	}
}
