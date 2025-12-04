package service

import (
	"app/pkg/dto"
	"database/sql"
	"errors"
	"testing"
)

func TestProjectService_CreateAndFindById(t *testing.T) {
	// Arrange
	svc, _ := setupProjectService(t)

	// Act
	created, err := svc.CreateProject(t.Context(), dto.CreateProjectDto{Slug: "proj-x", Name: "Project X"})
	if err != nil {
		t.Fatalf("CreateProject error: %v", err)
	}

	got, err := svc.FindProjectById(t.Context(), &created.ID)
	if err != nil {
		t.Fatalf("FindProjectById error: %v", err)
	}

	// Assert
	if got.ID != created.ID || got.Slug != "proj-x" || got.Name != "Project X" {
		t.Errorf("unexpected project: got=%+v", got)
	}
}

func TestProjectService_FindAllProjects(t *testing.T) {
	// Arrange
	svc, _ := setupProjectService(t)
	// seed a couple of projects via service and/or helper
	if _, err := svc.CreateProject(t.Context(), dto.CreateProjectDto{Slug: "a", Name: "A"}); err != nil {
		t.Fatalf("seed create A: %v", err)
	}
	if _, err := svc.CreateProject(t.Context(), dto.CreateProjectDto{Slug: "b", Name: "B"}); err != nil {
		t.Fatalf("seed create B: %v", err)
	}

	// Act
	got, err := svc.FindAllProjects(t.Context())

	// Assert
	if err != nil {
		t.Fatalf("FindAllProjects error: %v", err)
	}
	if len(got) < 2 {
		t.Fatalf("expected at least 2 projects, got %d", len(got))
	}
}

func TestProjectService_UpdateProject(t *testing.T) {
	// Arrange
	svc, _ := setupProjectService(t)
	created, err := svc.CreateProject(t.Context(), dto.CreateProjectDto{Slug: "p1", Name: "P1"})
	if err != nil {
		t.Fatalf("CreateProject error: %v", err)
	}

	// Act
	updated, err := svc.UpdateProject(t.Context(), &created.ID, dto.UpdateProjectDto{Slug: "p1-renamed", Name: "Project One"})

	// Assert
	if err != nil {
		t.Fatalf("UpdateProject error: %v", err)
	}
	if updated.Slug != "p1-renamed" || updated.Name != "Project One" {
		t.Errorf("update not applied: %+v", updated)
	}
}

func TestProjectService_DeleteProject(t *testing.T) {
	// Arrange
	svc, _ := setupProjectService(t)
	created, err := svc.CreateProject(t.Context(), dto.CreateProjectDto{Slug: "del", Name: "Delete Me"})
	if err != nil {
		t.Fatalf("CreateProject error: %v", err)
	}

	// Act
	if err := svc.DeleteProject(t.Context(), &created.ID); err != nil {
		t.Fatalf("DeleteProject error: %v", err)
	}

	// Assert: should not be found
	if _, err := svc.FindProjectById(t.Context(), &created.ID); err == nil {
		t.Fatalf("expected error on Get after delete")
	}
}

func TestProjectService_FindProjectById_NotFound(t *testing.T) {
	// Arrange
	svc, _ := setupProjectService(t)
	id := int64(999999)

	// Act
	_, err := svc.FindProjectById(t.Context(), &id)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected sql.ErrNoRows, got %v", err)
	}
}
