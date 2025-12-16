package controller

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/paginate"
	"app/pkg/service"
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mocks & helpers ---

type mockFileService struct{ mock.Mock }

func (m *mockFileService) FindAllFiles(ctx context.Context, params *dto.FindAllFilesParams) (*dto.FindAllFilesResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.FindAllFilesResult), args.Error(1)
}
func (m *mockFileService) FindFileById(ctx context.Context, params *dto.FindFileByIdParams) (*dto.FindFileByIdResult, error) {
	args := m.Called(ctx, params.ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.FindFileByIdResult), args.Error(1)
}
func (m *mockFileService) CreateFile(ctx context.Context, params *dto.CreateFileParams) (*dto.CreateFileResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CreateFileResult), args.Error(1)
}
func (m *mockFileService) UploadFile(ctx context.Context, params *dto.UploadFileParams) (*dto.UploadFileResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UploadFileResult), args.Error(1)
}
func (m *mockFileService) DownloadFile(ctx context.Context, params *dto.DownloadFileParams) (*dto.DownloadFileResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.DownloadFileResult), args.Error(1)
}
func (m *mockFileService) DeleteFile(ctx context.Context, params *dto.DeleteFileParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

type readSeekCloser struct{ *bytes.Reader }

func (r *readSeekCloser) Close() error { return nil }

func setupFileRouter(svc service.FileService) *chi.Mux {
	c := NewFileController(svc)
	r := chi.NewRouter()
	r.Route("/files", func(r chi.Router) {
		r.With(paginate.Paginate).Get("/", c.FindAllFiles)
		r.Post("/", c.CreateFile)
		r.Route("/{fileId}", func(r chi.Router) {
			r.Use(c.FileCtx)
			r.Get("/", c.GetFile)
			r.Post("/upload", c.UploadFile)
			r.Get("/download", c.DownloadFile)
			r.Delete("/", c.DeleteFile)
		})
	})
	return r
}

// --- Tests ---

func TestFileController_FindAllFiles(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	items := []*database.File{{ID: 1, Name: "a"}}
	svc.On("FindAllFiles", mock.Anything, mock.AnythingOfType("*dto.FindAllFilesParams")).Return(
		&dto.FindAllFilesResult{Files: items, Total: 1, Limit: 5, Offset: 0}, nil,
	)

	// Act
	req, _ := http.NewRequest("GET", "/files?versionId=10&limit=5&offset=0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, float64(1), body["total"])
	svc.AssertExpectations(t)
}

func TestFileController_FindAllFiles_BadQuery(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)

	// Act
	req, _ := http.NewRequest("GET", "/files", nil) // missing versionId
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFileController_GetFile(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	svc.On("FindFileById", mock.Anything, int64(3)).Return(&dto.FindFileByIdResult{File: &database.File{ID: 3, Name: "f"}}, nil)

	// Act
	req, _ := http.NewRequest("GET", "/files/3", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFileController_CreateFile(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	svc.On("CreateFile", mock.Anything, mock.AnythingOfType("*dto.CreateFileParams")).Return(
		&dto.CreateFileResult{File: &database.File{ID: 5, Name: "n"}}, nil,
	)

	// Act
	req, _ := http.NewRequest("POST", "/files", strings.NewReader(`{"name":"n"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFileController_DeleteFile(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	svc.On("FindFileById", mock.Anything, int64(7)).Return(&dto.FindFileByIdResult{File: &database.File{ID: 7}}, nil)
	svc.On("DeleteFile", mock.Anything, mock.AnythingOfType("*dto.DeleteFileParams")).Return(nil)

	// Act
	req, _ := http.NewRequest("DELETE", "/files/7", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestFileController_Upload_Success(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	fw, _ := mw.CreateFormFile("file", "hello.txt")
	fw.Write([]byte("hello"))
	mw.Close()
	svc.On("FindFileById", mock.Anything, int64(10)).Return(&dto.FindFileByIdResult{File: &database.File{ID: 10, Name: "x"}}, nil).Once()
	svc.On("UploadFile", mock.Anything, mock.AnythingOfType("*dto.UploadFileParams")).Return(
		&dto.UploadFileResult{File: &database.File{ID: 10, Name: "x", IsComplete: true}}, nil,
	).Once()

	// Act
	req, _ := http.NewRequest("POST", "/files/10/upload", &b)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFileController_Upload_Conflict(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	fw, _ := mw.CreateFormFile("file", "hello.txt")
	fw.Write([]byte("hello"))
	mw.Close()
	svc.On("FindFileById", mock.Anything, int64(11)).Return(&dto.FindFileByIdResult{File: &database.File{ID: 11, Name: "x"}}, nil).Once()
	svc.On("UploadFile", mock.Anything, mock.AnythingOfType("*dto.UploadFileParams")).Return(nil, service.ErrFileAlreadyExists).Once()

	// Act
	req, _ := http.NewRequest("POST", "/files/11/upload", &b)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestFileController_Download_Success(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	svc.On("FindFileById", mock.Anything, int64(20)).Return(&dto.FindFileByIdResult{File: &database.File{ID: 20, Name: "d.bin"}}, nil).Once()
	reader := &readSeekCloser{bytes.NewReader([]byte("content"))}
	svc.On("DownloadFile", mock.Anything, mock.AnythingOfType("*dto.DownloadFileParams")).Return(
		&dto.DownloadFileResult{File: &database.File{ID: 20, Name: "d.bin"}, Reader: reader}, nil,
	).Once()

	// Act
	req, _ := http.NewRequest("GET", "/files/20/download", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFileController_Download_IncompleteFile_BadRequest(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	svc.On("FindFileById", mock.Anything, int64(21)).Return(&dto.FindFileByIdResult{File: &database.File{ID: 21, Name: "d.bin"}}, nil).Once()
	svc.On("DownloadFile", mock.Anything, mock.AnythingOfType("*dto.DownloadFileParams")).Return(nil, service.ErrIncompleteFile).Once()

	// Act
	req, _ := http.NewRequest("GET", "/files/21/download", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFileController_FileCtx_NotFound(t *testing.T) {
	// Arrange
	svc := &mockFileService{}
	r := setupFileRouter(svc)
	svc.On("FindFileById", mock.Anything, int64(404)).Return(nil, apperrors.ErrNotFound)

	// Act
	req, _ := http.NewRequest("GET", "/files/404", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}
