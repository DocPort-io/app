package controller

import (
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/paginate"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
		r.With(paginate.Paginate).Get("/", versionController.FindAllVersions)
		r.Post("/", versionController.CreateVersion)

		r.Route("/{versionId}", func(r chi.Router) {
			r.Use(versionController.VersionCtx)
			r.Get("/", versionController.GetVersion)
			r.Put("/", versionController.UpdateVersion)
			r.Delete("/", versionController.DeleteVersion)
		})
	})

	return r, versionService
}

// test helpers
func newJSONRequest(t *testing.T, method, target, body string) *http.Request {
	t.Helper()
	req, err := http.NewRequest(method, target, strings.NewReader(body))
	require.NoError(t, err)
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	return req
}

func decodeJSON(t *testing.T, buf *bytes.Buffer, v any) {
	t.Helper()
	require.NoError(t, json.Unmarshal(buf.Bytes(), v))
}

func TestVersionController_VersionCtx(t *testing.T) {
	// Arrange
	r, versionService := setupFixtures()
	versionService.On("FindVersionById", mock.Anything, mock.AnythingOfType("int64")).Return(&dto.FindVersionByIdResult{Version: &database.Version{ID: 123}}, nil)

	// Act
	req := newJSONRequest(t, http.MethodGet, "/123", "")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]any
	decodeJSON(t, w.Body, &response)
	assert.Equal(t, float64(123), response["id"])

	versionService.AssertExpectations(t)
}

func TestVersionController_GetVersion(t *testing.T) {
	// Arrange
	r, versionService := setupFixtures()
	versionService.On("FindVersionById", mock.Anything, mock.AnythingOfType("int64")).Return(&dto.FindVersionByIdResult{Version: &database.Version{ID: 123}}, nil)

	// Act
	req := newJSONRequest(t, http.MethodGet, "/123", "")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]any
	decodeJSON(t, w.Body, &response)
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
	req := newJSONRequest(t, http.MethodGet, "/?projectId=123", "")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// decode into a lightweight shape for clarity
	type listResp struct {
		Versions []struct {
			ID int64 `json:"id"`
		} `json:"versions"`
		Total int64 `json:"total"`
	}
	var resp listResp
	decodeJSON(t, w.Body, &resp)
	require.Len(t, resp.Versions, 1)
	assert.Equal(t, int64(123), resp.Versions[0].ID)
	assert.Equal(t, int64(1), resp.Total)

	versionService.AssertExpectations(t)
}

func TestVersionController_CreateVersion(t *testing.T) {
	// Arrange
	r, svc := setupFixtures()
	svc.On("CreateVersion", mock.Anything, mock.AnythingOfType("*dto.CreateVersionParams")).Return(
		&dto.CreateVersionResult{Version: &database.Version{ID: 11, ProjectID: 5, Name: "v"}}, nil,
	)

	// Act
	req := newJSONRequest(t, http.MethodPost, "/", `{"name":"v","projectId":5}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestVersionController_UpdateVersion(t *testing.T) {
	// Arrange
	r, svc := setupFixtures()
	svc.On("FindVersionById", mock.Anything, int64(7)).Return(
		&dto.FindVersionByIdResult{Version: &database.Version{ID: 7, Name: "old"}}, nil,
	)
	svc.On("UpdateVersion", mock.Anything, mock.AnythingOfType("*dto.UpdateVersionParams")).Return(
		&dto.UpdateVersionResult{Version: &database.Version{ID: 7, Name: "new"}}, nil,
	)

	// Act
	req := newJSONRequest(t, http.MethodPut, "/7", `{"name":"new"}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestVersionController_DeleteVersion(t *testing.T) {
	// Arrange
	r, svc := setupFixtures()
	svc.On("FindVersionById", mock.Anything, int64(8)).Return(
		&dto.FindVersionByIdResult{Version: &database.Version{ID: 8}}, nil,
	)
	svc.On("DeleteVersion", mock.Anything, mock.AnythingOfType("*dto.DeleteVersionParams")).Return(nil)

	// Act
	req := newJSONRequest(t, http.MethodDelete, "/8", "")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Code)
}
