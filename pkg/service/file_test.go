package service

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/paginate"
	"app/pkg/storage"
	"bytes"
	"context"
	"database/sql"
	"io"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

// memStorage is a simple in-memory implementation of storage.FileStorage for tests.
type memStorage struct{ files map[string][]byte }

func newMemStorage() *memStorage { return &memStorage{files: map[string][]byte{}} }

func (m *memStorage) Save(_ context.Context, relativePath string, data io.Reader) error {
	b, err := io.ReadAll(data)
	if err != nil {
		return err
	}
	m.files[relativePath] = b
	return nil
}

type readSeekNopCloser struct{ *bytes.Reader }

func (r *readSeekNopCloser) Close() error { return nil }

func (m *memStorage) Retrieve(_ context.Context, relativePath string) (io.ReadSeekCloser, error) {
	b, ok := m.files[relativePath]
	if !ok {
		return nil, io.EOF
	}
	return &readSeekNopCloser{Reader: bytes.NewReader(b)}, nil
}

func (m *memStorage) Delete(_ context.Context, relativePath string) error {
	delete(m.files, relativePath)
	return nil
}

// Unused in tests, but required by interface; simple implementations.
func (m *memStorage) List(_ context.Context, _ string) ([]storage.ObjectInfo, error) { return nil, nil }
func (m *memStorage) Walk(_ context.Context, _ string, _ storage.WalkFunc) error     { return nil }

// Test: CreateFile creates a new DB record with default fields.
func TestFileService_CreateFile_CreatesRecord(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewFileService(queries, newMemStorage())

	res, err := svc.CreateFile(context.Background(), &dto.CreateFileParams{Name: "file.pdf"})
	assert.NoError(t, err)
	assert.NotZero(t, res.File.ID)
	assert.Equal(t, "file.pdf", res.File.Name)
	assert.False(t, res.File.IsComplete)
}

// Test: UploadFile writes content to storage, sets metadata and marks complete.
func TestFileService_UploadFile_SetsMetadataAndComplete(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	st := newMemStorage()
	svc := NewFileService(queries, st)

	created, err := svc.CreateFile(context.Background(), &dto.CreateFileParams{Name: "file.pdf"})
	assert.NoError(t, err)

	content := []byte("hello world")
	upFile := &readSeekNopCloser{Reader: bytes.NewReader(content)}
	fh := &multipart.FileHeader{Filename: "file.pdf", Size: int64(len(content))}
	upRes, err := svc.UploadFile(context.Background(), &dto.UploadFileParams{ID: created.File.ID, File: upFile, FileHeader: fh})
	assert.NoError(t, err)
	assert.NotNil(t, upRes.File.Path)
	assert.True(t, upRes.File.IsComplete)
	// MimeType may vary; ensure it is set when uploaded
	assert.NotNil(t, upRes.File.MimeType)
}

// Test: UploadFile on already complete file returns ErrFileAlreadyExists.
func TestFileService_UploadFile_OnAlreadyComplete_ReturnsErr(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	st := newMemStorage()
	svc := NewFileService(queries, st)

	created, err := svc.CreateFile(context.Background(), &dto.CreateFileParams{Name: "file.pdf"})
	assert.NoError(t, err)

	content := []byte("hello world")
	upFile1 := &readSeekNopCloser{Reader: bytes.NewReader(content)}
	fh := &multipart.FileHeader{Filename: "file.pdf", Size: int64(len(content))}
	_, err = svc.UploadFile(context.Background(), &dto.UploadFileParams{ID: created.File.ID, File: upFile1, FileHeader: fh})
	assert.NoError(t, err)

	// Second upload should fail with ErrFileAlreadyExists
	upFile2 := &readSeekNopCloser{Reader: bytes.NewReader(content)}
	_, err = svc.UploadFile(context.Background(), &dto.UploadFileParams{ID: created.File.ID, File: upFile2, FileHeader: fh})
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrFileAlreadyExists)
}

// Test: DownloadFile returns uploaded content.
func TestFileService_DownloadFile_ReturnsContent(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	st := newMemStorage()
	svc := NewFileService(queries, st)

	created, err := svc.CreateFile(context.Background(), &dto.CreateFileParams{Name: "file.pdf"})
	assert.NoError(t, err)
	content := []byte("hello world")
	upFile := &readSeekNopCloser{Reader: bytes.NewReader(content)}
	fh := &multipart.FileHeader{Filename: "file.pdf", Size: int64(len(content))}
	_, err = svc.UploadFile(context.Background(), &dto.UploadFileParams{ID: created.File.ID, File: upFile, FileHeader: fh})
	assert.NoError(t, err)

	dl, err := svc.DownloadFile(context.Background(), &dto.DownloadFileParams{ID: created.File.ID})
	assert.NoError(t, err)
	got, _ := io.ReadAll(dl.Reader)
	assert.Equal(t, content, got)
}

// Test: DownloadFile on incomplete file returns ErrIncompleteFile.
func TestFileService_DownloadFile_Incomplete_ReturnsErr(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewFileService(queries, newMemStorage())

	created, err := svc.CreateFile(context.Background(), &dto.CreateFileParams{Name: "file.pdf"})
	assert.NoError(t, err)

	_, err = svc.DownloadFile(context.Background(), &dto.DownloadFileParams{ID: created.File.ID})
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrIncompleteFile)
}

// Test: DeleteFile removes DB row and underlying storage object.
func TestFileService_DeleteFile_RemovesDBAndStorage(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	st := newMemStorage()
	svc := NewFileService(queries, st)

	created, err := svc.CreateFile(context.Background(), &dto.CreateFileParams{Name: "file.pdf"})
	assert.NoError(t, err)
	content := []byte("bytes")
	upFile := &readSeekNopCloser{Reader: bytes.NewReader(content)}
	fh := &multipart.FileHeader{Filename: "file.pdf", Size: int64(len(content))}
	upRes, err := svc.UploadFile(context.Background(), &dto.UploadFileParams{ID: created.File.ID, File: upFile, FileHeader: fh})
	assert.NoError(t, err)
	path := *upRes.File.Path

	// Act
	err = svc.DeleteFile(context.Background(), &dto.DeleteFileParams{ID: created.File.ID})
	assert.NoError(t, err)

	// Assert storage object removed
	_, err = st.Retrieve(context.Background(), path)
	assert.Error(t, err)

	// Assert DB row removed
	_, err = queries.GetFile(context.Background(), created.File.ID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}

// Test: FindAllFiles paginates results for a version.
func TestFileService_FindAllFiles_Paginates(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	st := newMemStorage()
	svc := NewFileService(queries, st)

	// Arrange: project, version, two files attached
	proj, err := queries.CreateProject(context.Background(), &database.CreateProjectParams{Slug: "p", Name: "P"})
	assert.NoError(t, err)
	ver, err := queries.CreateVersion(context.Background(), &database.CreateVersionParams{Name: "v1", ProjectID: proj.ID})
	assert.NoError(t, err)

	mk := func(name string) int64 {
		r, err := svc.CreateFile(context.Background(), &dto.CreateFileParams{Name: name})
		assert.NoError(t, err)
		content := []byte(name)
		up := &readSeekNopCloser{Reader: bytes.NewReader(content)}
		fh := &multipart.FileHeader{Filename: name, Size: int64(len(content))}
		_, err = svc.UploadFile(context.Background(), &dto.UploadFileParams{ID: r.File.ID, File: up, FileHeader: fh})
		assert.NoError(t, err)
		err = queries.AttachFileToVersion(context.Background(), &database.AttachFileToVersionParams{VersionID: ver.ID, FileID: r.File.ID})
		assert.NoError(t, err)
		return r.File.ID
	}
	_ = mk("a.pdf")
	_ = mk("b.pdf")

	// Act
	res, err := svc.FindAllFiles(context.Background(), &dto.FindAllFilesParams{VersionID: ver.ID, Pagination: &paginate.Pagination{Limit: 1, Offset: 0}})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int64(2), res.Total)
	assert.Len(t, res.Files, 1)
}

// Test: FindFileById returns not found error for missing file.
func TestFileService_FindFileById_NotFound_ReturnsErr(t *testing.T) {
	_, queries := newTestDBAndQueries(t)
	svc := NewFileService(queries, newMemStorage())

	_, err := svc.FindFileById(context.Background(), &dto.FindFileByIdParams{ID: 42})
	assert.Error(t, err)
	assert.ErrorIs(t, err, apperrors.ErrNotFound)
}
