package controller

import (
	"app/pkg/database"
	"app/pkg/dto"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockVersionService struct {
	mock.Mock
}

func (m *mockVersionService) FindAllVersions(ctx context.Context, params *dto.FindAllVersionsParams) (*dto.FindAllVersionsResult, error) {
	args := m.Called(ctx, params.ProjectID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.FindAllVersionsResult), args.Error(1)
}

func (m *mockVersionService) FindVersionById(ctx context.Context, params *dto.FindVersionByIdParams) (*dto.FindVersionByIdResult, error) {
	args := m.Called(ctx, params.ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.FindVersionByIdResult), args.Error(1)
}

func (m *mockVersionService) CreateVersion(ctx context.Context, params *dto.CreateVersionParams) (*dto.CreateVersionResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CreateVersionResult), args.Error(1)
}

func (m *mockVersionService) UpdateVersion(ctx context.Context, params *dto.UpdateVersionParams) (*dto.UpdateVersionResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UpdateVersionResult), args.Error(1)
}

func (m *mockVersionService) DeleteVersion(ctx context.Context, params *dto.DeleteVersionParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *mockVersionService) AttachFileToVersion(ctx context.Context, params *dto.AttachFileToVersionParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *mockVersionService) DetachFileFromVersion(ctx context.Context, params *dto.DetachFileFromVersionParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func setupFixtures() (*chi.Mux, *mockVersionService) {
	versionService := &mockVersionService{}
	versionController := NewVersionController(versionService)

	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", versionController.FindAllVersions)
		r.Post("/", versionController.CreateVersion)

		r.Route("/{versionId}", func(r chi.Router) {
			r.Use(versionController.VersionCtx)
			r.Get("/", versionController.GetVersion)
			r.Post("/", versionController.UpdateVersion)
			r.Delete("/", versionController.DeleteVersion)
		})
	})

	return r, versionService
}

func TestVersionController_VersionCtx(t *testing.T) {
	// Arrange
	r, versionService := setupFixtures()
	versionService.On("FindVersionById", mock.Anything, mock.AnythingOfType("int64")).Return(&dto.FindVersionByIdResult{Version: &database.Version{ID: 123}}, nil)

	// Act
	req, _ := http.NewRequest("GET", "/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(123), response["id"])

	versionService.AssertExpectations(t)
}

func TestVersionController_GetVersion(t *testing.T) {
	// Arrange
	r, versionService := setupFixtures()
	versionService.On("FindVersionById", mock.Anything, mock.AnythingOfType("int64")).Return(&dto.FindVersionByIdResult{Version: &database.Version{ID: 123}}, nil)

	// Act
	req, _ := http.NewRequest("GET", "/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(123), response["id"])

	versionService.AssertExpectations(t)
}

func TestVersionController_FindAllVersions(t *testing.T) {
	// Arrange
	r, versionService := setupFixtures()
	items := make([]*database.Version, 1)
	items[0] = &database.Version{ID: 123}
	versionService.On("FindAllVersions", mock.Anything, mock.AnythingOfType("int64")).Return(&dto.FindAllVersionsResult{Versions: items, Total: int64(1)}, nil)

	// Act
	req, _ := http.NewRequest("GET", "/?projectId=123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	//var response map[string]interface{}
	//json.Unmarshal(w.Body.Bytes(), &response)
	//assert.Equal(t, int64(123), response["versions"][0]["id"])

	versionService.AssertExpectations(t)
}
