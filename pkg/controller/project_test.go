package controller

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/paginate"
	"app/pkg/service"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockProjectService struct{ mock.Mock }

func (m *mockProjectService) FindAllProjects(ctx context.Context, params *dto.FindAllProjectsParams) (*dto.FindAllProjectsResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.FindAllProjectsResult), args.Error(1)
}

func (m *mockProjectService) FindProjectById(ctx context.Context, params *dto.FindProjectByIdParams) (*dto.FindProjectByIdResult, error) {
	args := m.Called(ctx, params.ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.FindProjectByIdResult), args.Error(1)
}

func (m *mockProjectService) CreateProject(ctx context.Context, params *dto.CreateProjectParams) (*dto.CreateProjectResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CreateProjectResult), args.Error(1)
}

func (m *mockProjectService) UpdateProject(ctx context.Context, params *dto.UpdateProjectParams) (*dto.UpdateProjectResult, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UpdateProjectResult), args.Error(1)
}

func (m *mockProjectService) DeleteProject(ctx context.Context, params *dto.DeleteProjectParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func setupProjectRouter(svc service.ProjectService) *chi.Mux {
	c := NewProjectController(svc)
	r := chi.NewRouter()
	r.Route("/projects", func(r chi.Router) {
		r.With(paginate.Paginate).Get("/", c.FindAllProjects)
		r.Post("/", c.CreateProject)
		r.Route("/{projectId}", func(r chi.Router) {
			r.Use(c.ProjectCtx)
			r.Get("/", c.GetProject)
			r.Put("/", c.UpdateProject)
			r.Delete("/", c.DeleteProject)
		})
	})
	return r
}

func TestProjectController_FindAllProjects(t *testing.T) {
	// Arrange
	svc := &mockProjectService{}
	r := setupProjectRouter(svc)
	items := []*database.ListProjectsWithLocationsRow{{ID: 1}}
	svc.On("FindAllProjects", mock.Anything, mock.AnythingOfType("*dto.FindAllProjectsParams")).Return(
		&dto.FindAllProjectsResult{Projects: items, Total: 1, Limit: 10, Offset: 0}, nil,
	)

	// Act
	req, _ := http.NewRequest("GET", "/projects?limit=10&offset=0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, float64(1), body["total"])
	svc.AssertExpectations(t)
}

func TestProjectController_ProjectCtx_and_GetProject(t *testing.T) {
	// Arrange
	svc := &mockProjectService{}
	r := setupProjectRouter(svc)
	svc.On("FindProjectById", mock.Anything, int64(123)).Return(
		&dto.FindProjectByIdResult{Project: &database.Project{ID: 123, Slug: "p", Name: "Proj"}}, nil,
	)

	// Act
	req, _ := http.NewRequest("GET", "/projects/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, float64(123), body["id"])
	svc.AssertExpectations(t)
}

func TestProjectController_CreateProject(t *testing.T) {
	// Arrange
	svc := &mockProjectService{}
	r := setupProjectRouter(svc)
	svc.On("CreateProject", mock.Anything, mock.AnythingOfType("*dto.CreateProjectParams")).Return(
		&dto.CreateProjectResult{Project: &database.Project{ID: 42, Slug: "s", Name: "N"}}, nil,
	)

	// Act
	body := `{"slug":"s","name":"N"}`
	req, _ := http.NewRequest("POST", "/projects", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, float64(42), resp["id"])
	svc.AssertExpectations(t)
}

func TestProjectController_UpdateProject(t *testing.T) {
	// Arrange
	svc := &mockProjectService{}
	r := setupProjectRouter(svc)
	svc.On("FindProjectById", mock.Anything, int64(7)).Return(
		&dto.FindProjectByIdResult{Project: &database.Project{ID: 7, Slug: "old", Name: "old"}}, nil,
	)
	svc.On("UpdateProject", mock.Anything, mock.AnythingOfType("*dto.UpdateProjectParams")).Return(
		&dto.UpdateProjectResult{Project: &database.Project{ID: 7, Slug: "new", Name: "new"}}, nil,
	)

	// Act
	body := `{"slug":"new","name":"new"}`
	req, _ := http.NewRequest("PUT", "/projects/7", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestProjectController_DeleteProject(t *testing.T) {
	// Arrange
	svc := &mockProjectService{}
	r := setupProjectRouter(svc)
	svc.On("FindProjectById", mock.Anything, int64(9)).Return(&dto.FindProjectByIdResult{Project: &database.Project{ID: 9}}, nil)
	svc.On("DeleteProject", mock.Anything, mock.AnythingOfType("*dto.DeleteProjectParams")).Return(nil)

	// Act
	req, _ := http.NewRequest("DELETE", "/projects/9", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Code)
	svc.AssertExpectations(t)
}

func TestProjectController_ProjectCtx_NotFound(t *testing.T) {
	// Arrange
	svc := &mockProjectService{}
	r := setupProjectRouter(svc)
	svc.On("FindProjectById", mock.Anything, int64(404)).Return(nil, apperrors.ErrNotFound)

	// Act
	req, _ := http.NewRequest("GET", "/projects/404", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestProjectController_ProjectCtx_InternalServerError(t *testing.T) {
	// Arrange
	svc := &mockProjectService{}
	r := setupProjectRouter(svc)
	svc.On("FindProjectById", mock.Anything, int64(500)).Return(nil, errors.New("boom"))

	// Act
	req, _ := http.NewRequest("GET", "/projects/500", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
