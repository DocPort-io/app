package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewFilesystemStorage(t *testing.T) {
	t.Run("creates root directory if not exists", func(t *testing.T) {
		tmpDir := t.TempDir()
		rootPath := filepath.Join(tmpDir, "storage", "nested", "path")

		storage, err := NewFilesystemStorage(rootPath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if storage == nil {
			t.Fatal("expected storage to be non-nil")
		}
		defer func() {
			if fs, ok := storage.(*filesystemStorage); ok {
				_ = fs.root.Close()
			}
		}()

		// Check directory was created
		if _, err := os.Stat(rootPath); os.IsNotExist(err) {
			t.Fatal("expected root directory to be created")
		}
	})

	t.Run("opens existing directory", func(t *testing.T) {
		tmpDir := t.TempDir()

		storage, err := NewFilesystemStorage(tmpDir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if storage == nil {
			t.Fatal("expected storage to be non-nil")
		}
		defer func() {
			if fs, ok := storage.(*filesystemStorage); ok {
				_ = fs.root.Close()
			}
		}()
	})
}

func TestFilesystemStorage_Save(t *testing.T) {
	t.Run("saves file successfully", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()
		content := []byte("test content")
		relativePath := "test-file.txt"

		err := storage.Save(ctx, relativePath, bytes.NewReader(content))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify file exists and has correct content
		retrieved, err := storage.Retrieve(ctx, relativePath)
		if err != nil {
			t.Fatalf("failed to retrieve file: %v", err)
		}
		defer retrieved.Close()

		gotContent, err := io.ReadAll(retrieved)
		if err != nil {
			t.Fatalf("failed to read content: %v", err)
		}

		if !bytes.Equal(gotContent, content) {
			t.Errorf("expected content %q, got %q", content, gotContent)
		}
	})

	t.Run("saves file in nested directory", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()
		content := []byte("nested content")
		relativePath := "subdir/nested/file.txt"

		err := storage.Save(ctx, relativePath, bytes.NewReader(content))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify file exists
		retrieved, err := storage.Retrieve(ctx, relativePath)
		if err != nil {
			t.Fatalf("failed to retrieve file: %v", err)
		}
		defer retrieved.Close()
	})

	t.Run("overwrites existing file", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()
		relativePath := "overwrite-test.txt"

		// Save first version
		err := storage.Save(ctx, relativePath, bytes.NewReader([]byte("first")))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Overwrite with second version
		err = storage.Save(ctx, relativePath, bytes.NewReader([]byte("second")))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify file has second content
		retrieved, err := storage.Retrieve(ctx, relativePath)
		if err != nil {
			t.Fatalf("failed to retrieve file: %v", err)
		}
		defer retrieved.Close()

		gotContent, err := io.ReadAll(retrieved)
		if err != nil {
			t.Fatalf("failed to read content: %v", err)
		}

		if string(gotContent) != "second" {
			t.Errorf("expected content %q, got %q", "second", gotContent)
		}
	})

	t.Run("handles forward slashes in path", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()
		content := []byte("slash test")
		relativePath := "path/with/forward/slashes.txt"

		err := storage.Save(ctx, relativePath, bytes.NewReader(content))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify file can be retrieved
		retrieved, err := storage.Retrieve(ctx, relativePath)
		if err != nil {
			t.Fatalf("failed to retrieve file: %v", err)
		}
		defer retrieved.Close()
	})
}

func TestFilesystemStorage_Retrieve(t *testing.T) {
	t.Run("retrieves existing file", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()
		content := []byte("retrieve test")
		relativePath := "retrieve-test.txt"

		// Save file first
		err := storage.Save(ctx, relativePath, bytes.NewReader(content))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		// Retrieve file
		retrieved, err := storage.Retrieve(ctx, relativePath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer retrieved.Close()

		gotContent, err := io.ReadAll(retrieved)
		if err != nil {
			t.Fatalf("failed to read content: %v", err)
		}

		if !bytes.Equal(gotContent, content) {
			t.Errorf("expected content %q, got %q", content, gotContent)
		}
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		_, err := storage.Retrieve(ctx, "non-existent.txt")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("supports seeking", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()
		content := []byte("0123456789")
		relativePath := "seek-test.txt"

		// Save file
		err := storage.Save(ctx, relativePath, bytes.NewReader(content))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		// Retrieve and seek
		retrieved, err := storage.Retrieve(ctx, relativePath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer retrieved.Close()

		// Seek to position 5
		_, err = retrieved.Seek(5, io.SeekStart)
		if err != nil {
			t.Fatalf("failed to seek: %v", err)
		}

		// Read remaining content
		gotContent, err := io.ReadAll(retrieved)
		if err != nil {
			t.Fatalf("failed to read content: %v", err)
		}

		expected := []byte("56789")
		if !bytes.Equal(gotContent, expected) {
			t.Errorf("expected content %q, got %q", expected, gotContent)
		}
	})
}

func TestFilesystemStorage_Delete(t *testing.T) {
	t.Run("deletes existing file", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()
		relativePath := "delete-test.txt"

		// Save file first
		err := storage.Save(ctx, relativePath, bytes.NewReader([]byte("delete me")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		// Delete file
		err = storage.Delete(ctx, relativePath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify file no longer exists
		_, err = storage.Retrieve(ctx, relativePath)
		if err == nil {
			t.Fatal("expected error when retrieving deleted file, got nil")
		}
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		err := storage.Delete(ctx, "non-existent.txt")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestFilesystemStorage_List(t *testing.T) {
	t.Run("lists files in directory", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		// Create test files
		files := []struct {
			path    string
			content []byte
		}{
			{"file1.txt", []byte("content1")},
			{"file2.txt", []byte("content2")},
			{"file3.txt", []byte("longer content 3")},
		}

		for _, f := range files {
			err := storage.Save(ctx, f.path, bytes.NewReader(f.content))
			if err != nil {
				t.Fatalf("failed to save file %s: %v", f.path, err)
			}
		}

		// List files
		objects, err := storage.List(ctx, ".")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(objects) != len(files) {
			t.Fatalf("expected %d objects, got %d", len(files), len(objects))
		}

		// Verify each file is in the list
		for _, f := range files {
			found := false
			for _, obj := range objects {
				if strings.HasSuffix(obj.Path, f.path) {
					found = true
					if obj.Size != int64(len(f.content)) {
						t.Errorf("expected size %d for %s, got %d", len(f.content), f.path, obj.Size)
					}
					break
				}
			}
			if !found {
				t.Errorf("file %s not found in list", f.path)
			}
		}
	})

	t.Run("lists files in subdirectory", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		// Create files in subdirectory
		subdir := "subdir"
		err := storage.Save(ctx, filepath.Join(subdir, "file1.txt"), bytes.NewReader([]byte("content1")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		err = storage.Save(ctx, filepath.Join(subdir, "file2.txt"), bytes.NewReader([]byte("content2")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		// Create file in root (should not be listed)
		err = storage.Save(ctx, "root-file.txt", bytes.NewReader([]byte("root")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		// List files in subdirectory
		objects, err := storage.List(ctx, subdir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(objects) != 2 {
			t.Fatalf("expected 2 objects, got %d", len(objects))
		}
	})

	t.Run("returns empty list for empty directory", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		objects, err := storage.List(ctx, ".")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(objects) != 0 {
			t.Fatalf("expected 0 objects, got %d", len(objects))
		}
	})

	t.Run("does not list directories", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		// Create nested directories with files
		err := storage.Save(ctx, "dir1/file1.txt", bytes.NewReader([]byte("content1")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		err = storage.Save(ctx, "dir2/file2.txt", bytes.NewReader([]byte("content2")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		// List root - should not include directories
		objects, err := storage.List(ctx, ".")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(objects) != 0 {
			t.Fatalf("expected 0 objects (directories should be excluded), got %d", len(objects))
		}
	})
}

func TestFilesystemStorage_Walk(t *testing.T) {
	t.Run("walks all files recursively", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		// Create nested file structure
		files := []struct {
			path    string
			content []byte
		}{
			{"file1.txt", []byte("content1")},
			{"dir1/file2.txt", []byte("content2")},
			{"dir1/subdir/file3.txt", []byte("content3")},
			{"dir2/file4.txt", []byte("content4")},
		}

		for _, f := range files {
			err := storage.Save(ctx, f.path, bytes.NewReader(f.content))
			if err != nil {
				t.Fatalf("failed to save file %s: %v", f.path, err)
			}
		}

		// Walk and collect files
		var walked []ObjectInfo
		err := storage.Walk(ctx, ".", func(info ObjectInfo) error {
			walked = append(walked, info)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(walked) != len(files) {
			t.Fatalf("expected %d files, got %d", len(files), len(walked))
		}

		// Verify each file was walked
		for _, f := range files {
			found := false
			for _, obj := range walked {
				if strings.HasSuffix(obj.Path, f.path) {
					found = true
					if obj.Size != int64(len(f.content)) {
						t.Errorf("expected size %d for %s, got %d", len(f.content), f.path, obj.Size)
					}
					break
				}
			}
			if !found {
				t.Errorf("file %s not found in walk", f.path)
			}
		}
	})

	t.Run("walks subdirectory only", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		// Create files in different directories
		err := storage.Save(ctx, "root-file.txt", bytes.NewReader([]byte("root")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		err = storage.Save(ctx, "subdir/file1.txt", bytes.NewReader([]byte("content1")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		err = storage.Save(ctx, "subdir/nested/file2.txt", bytes.NewReader([]byte("content2")))
		if err != nil {
			t.Fatalf("failed to save file: %v", err)
		}

		// Walk only subdir
		var walked []ObjectInfo
		err = storage.Walk(ctx, "subdir", func(info ObjectInfo) error {
			walked = append(walked, info)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(walked) != 2 {
			t.Fatalf("expected 2 files, got %d", len(walked))
		}

		// Verify root file is not included
		for _, obj := range walked {
			if strings.Contains(obj.Path, "root-file.txt") {
				t.Error("root-file.txt should not be included in walk")
			}
		}
	})

	t.Run("stops walking on error", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		// Create multiple files
		for i := 1; i <= 5; i++ {
			path := filepath.Join(fmt.Sprintf("file%d.txt", i))
			err := storage.Save(ctx, path, bytes.NewReader([]byte("content")))
			if err != nil {
				t.Fatalf("failed to save file: %v", err)
			}
		}

		// Walk and stop after 2 files
		walked := 0
		expectedErr := io.EOF
		err := storage.Walk(ctx, ".", func(info ObjectInfo) error {
			walked++
			if walked == 2 {
				return expectedErr
			}
			return nil
		})

		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}

		if walked != 2 {
			t.Errorf("expected to walk 2 files, got %d", walked)
		}
	})

	t.Run("handles empty directory", func(t *testing.T) {
		storage, cleanup := setupStorage(t)
		defer cleanup()

		ctx := context.Background()

		walked := 0
		err := storage.Walk(ctx, ".", func(info ObjectInfo) error {
			walked++
			return nil
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if walked != 0 {
			t.Errorf("expected 0 files, got %d", walked)
		}
	})
}

// setupStorage creates a temporary filesystem storage for testing
func setupStorage(t *testing.T) (FileStorage, func()) {
	t.Helper()

	tmpDir := t.TempDir()
	storage, err := NewFilesystemStorage(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	return storage, func() {
		// Close the root to release file handles
		if fs, ok := storage.(*filesystemStorage); ok {
			_ = fs.root.Close()
		}
	}
}
