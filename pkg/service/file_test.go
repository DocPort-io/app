package service

import (
	"app/pkg/dto"
	"database/sql"
	"errors"
	"testing"
)

func TestFileService_FindAllFiles(t *testing.T) {
	// Arrange
	service, _, queries, _ := setupFileService(t)

	var testFileName = "test-file.pdf"
	_, err := queries.CreateFile(t.Context(), testFileName)
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
	service, _, _, _ := setupFileService(t)

	// Act
	file := &dto.CreateFileDto{
		Name: "test-file.pdf",
	}

	result, err := service.CreateFile(t.Context(), *file)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Errorf("expected a file, got nil")
	}
}

func TestFileService_FindFileById(t *testing.T) {
	// Arrange
	service, _, queries, _ := setupFileService(t)

	var testFileName = "test-file.pdf"
	file, err := queries.CreateFile(t.Context(), testFileName)
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
	if got.ID != file.ID || got.Name != testFileName {
		t.Errorf("returned file does not match: got %+v, want %+v", got, testFileName)
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
