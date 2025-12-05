package service

import (
	"app/pkg/dto"
	"database/sql"
	"errors"
	"testing"
)

func TestVersionService_CreateAndFindById(t *testing.T) {
	// Arrange
	vs, _, q, _, _ := setupVersionService(t)
	p := seedProject(t, q, "proj-a", "Project A")
	var desc = "initial"

	// Act
	created, err := vs.CreateVersion(t.Context(), dto.CreateVersionDto{Name: "v1", Description: &desc, ProjectId: p.ID})
	if err != nil {
		t.Fatalf("CreateVersion error: %v", err)
	}

	got, err := vs.FindVersionById(t.Context(), &created.ID)
	if err != nil {
		t.Fatalf("FindVersionById error: %v", err)
	}

	// Assert
	if got.ID != created.ID || got.Name != "v1" || got.ProjectID != p.ID {
		t.Errorf("unexpected version: got=%+v", got)
	}
	if got.Description == nil || *got.Description != desc {
		t.Errorf("description mismatch: got=%v want=%v", got.Description, desc)
	}
}

func TestVersionService_FindAllVersions(t *testing.T) {
	// Arrange
	vs, _, q, _, _ := setupVersionService(t)
	p := seedProject(t, q, "proj-b", "Project B")
	seedVersion(t, q, p.ID, "v1", nil)
	seedVersion(t, q, p.ID, "v2", nil)

	// Act
	got, err := vs.FindAllVersions(t.Context(), nil)

	// Assert
	if err != nil {
		t.Fatalf("FindAllVersions error: %v", err)
	}
	if len(got) < 2 {
		t.Fatalf("expected at least 2 versions, got %d", len(got))
	}
}

func TestVersionService_FindAllVersions_FilterByProject(t *testing.T) {
	// Arrange
	vs, _, q, _, _ := setupVersionService(t)
	p1 := seedProject(t, q, "proj-c1", "Project C1")
	p2 := seedProject(t, q, "proj-c2", "Project C2")
	v1 := seedVersion(t, q, p1.ID, "v1", nil)
	_ = v1
	seedVersion(t, q, p2.ID, "v2", nil)

	// Act
	got, err := vs.FindAllVersions(t.Context(), &p1.ID)

	// Assert
	if err != nil {
		t.Fatalf("FindAllVersions (filtered) error: %v", err)
	}
	if len(got) != 1 || got[0].ProjectID != p1.ID {
		t.Fatalf("expected exactly 1 version for project %d, got: %+v", p1.ID, got)
	}
}

func TestVersionService_UpdateVersion(t *testing.T) {
	// Arrange
	vs, _, q, _, _ := setupVersionService(t)
	p := seedProject(t, q, "proj-d", "Project D")
	var desc = "initial"
	v := seedVersion(t, q, p.ID, "v1", &desc)

	// Act
	var descNew = "second"
	updated, err := vs.UpdateVersion(t.Context(), &v.ID, dto.UpdateVersionDto{Name: "v1.1", Description: &descNew})

	// Assert
	if err != nil {
		t.Fatalf("UpdateVersion error: %v", err)
	}
	if updated.Name != "v1.1" || updated.Description == nil || *updated.Description != "second" {
		t.Errorf("update not applied: %+v", updated)
	}
}

func TestVersionService_DeleteVersion(t *testing.T) {
	// Arrange
	vs, _, q, _, _ := setupVersionService(t)
	p := seedProject(t, q, "proj-e", "Project E")
	v := seedVersion(t, q, p.ID, "v1", nil)

	// Act
	if err := vs.DeleteVersion(t.Context(), &v.ID); err != nil {
		t.Fatalf("DeleteVersion error: %v", err)
	}

	// Assert: should not be found
	if _, err := vs.FindVersionById(t.Context(), &v.ID); err == nil {
		t.Fatalf("expected error on Get after delete")
	}
}

func TestVersionService_FindVersionById_NotFound(t *testing.T) {
	// Arrange
	vs, _, _, _, _ := setupVersionService(t)
	id := int64(999999)

	// Act
	_, err := vs.FindVersionById(t.Context(), &id)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected sql.ErrNoRows, got %v", err)
	}
}

func TestVersionService_UploadFileToVersion(t *testing.T) {
	// Arrange
	vs, fileSvc, q, _, spy := setupVersionService(t)
	p := seedProject(t, q, "proj-f", "Project F")
	v := seedVersion(t, q, p.ID, "v1", nil)

	// Act
	tmp := createTempFile(t)
	uploaded, err := vs.UploadFileToVersion(t.Context(), &v.ID, dto.UploadFileToVersionDto{Name: "doc.pdf", Size: 13, Path: tmp})

	// Assert
	if err != nil {
		t.Fatalf("AttachFileToVersion error: %v", err)
	}
	if uploaded == nil || uploaded.ID == 0 {
		t.Fatalf("expected a created file, got %+v", uploaded)
	}

	// verify join via ListFilesByVersionId
	files, err := q.ListFilesByVersionId(t.Context(), v.ID)
	if err != nil {
		t.Fatalf("ListFilesByVersionId error: %v", err)
	}
	if len(files) != 1 || files[0].ID != uploaded.ID {
		t.Fatalf("expected file to be attached to version; got: %+v", files)
	}

	// verify that underlying FileService used storage.Save once
	_ = fileSvc
	if len(spy.Calls) == 0 || spy.Calls[0] != "Save" {
		t.Fatalf("expected storage Save to be called, got calls: %+v", spy.Calls)
	}
}
