package controller

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/httputil"
	"app/pkg/service"
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

const ProjectCtxKey = "project"

type ProjectController struct {
	projectService service.ProjectService
}

func NewProjectController(projectService service.ProjectService) *ProjectController {
	return &ProjectController{projectService: projectService}
}

func (c *ProjectController) ProjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projectId, err := httputil.URLParamInt64(r, "projectId")
		if err != nil {
			httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
			return
		}

		project, err := c.projectService.FindProjectById(r.Context(), projectId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNotFound) {
				httputil.Render(w, r, apperrors.ErrHTTPNotFoundError())
				return
			}

			httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
			return
		}

		ctx := context.WithValue(r.Context(), ProjectCtxKey, project)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getProject(ctx context.Context) *database.Project {
	return ctx.Value(ProjectCtxKey).(*database.Project)
}

// FindAllProjects godoc
//
//	@summary	Find all projects
//	@tags		projects
//	@accept		json
//	@produce	json
//	@success	200	{object}	dto.ListProjectsResponseDto
//	@failure	400	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/projects [get]
func (c *ProjectController) FindAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, total, err := c.projectService.FindAllProjects(r.Context())
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToListProjectsResponse(projects, total))
}

// GetProject godoc
//
//	@summary	Get a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		id	path		uint	true	"Project identifier"
//	@success	200	{object}	dto.ProjectResponseDto
//	@failure	400	{object}	apperrors.ErrResponse
//	@failure	404	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/projects/{projectId} [get]
func (c *ProjectController) GetProject(w http.ResponseWriter, r *http.Request) {
	project := getProject(r.Context())
	httputil.Render(w, r, dto.ToProjectResponse(project))
}

// CreateProject godoc
//
//	@summary	Create a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		request	body		dto.CreateProjectDto	true	"Create a project"
//	@success	201		{object}	dto.ProjectResponseDto
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/projects [post]
func (c *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	input := &dto.CreateProjectDto{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	project, err := c.projectService.CreateProject(r.Context(), input)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	httputil.Render(w, r, dto.ToProjectResponse(project))
}

// UpdateProject godoc
//
//	@summary	Update a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		id		path		uint					true	"Project identifier"
//	@param		request	body		dto.UpdateProjectDto	true	"Update a project"
//	@success	200		{object}	dto.ProjectResponseDto
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	404		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/projects/{projectId} [put]
func (c *ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	project := getProject(r.Context())

	input := &dto.UpdateProjectDto{
		Slug: project.Slug,
		Name: project.Name,
	}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	project, err := c.projectService.UpdateProject(r.Context(), project.ID, input)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToProjectResponse(project))
}

// DeleteProject godoc
//
//	@summary	Delete a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		id	path	uint	true	"Project identifier"
//	@success	204
//	@failure	404	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/projects/{projectId} [delete]
func (c *ProjectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	project := getProject(r.Context())

	err := c.projectService.DeleteProject(r.Context(), project.ID)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}
