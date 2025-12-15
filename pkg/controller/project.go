package controller

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/httputil"
	"app/pkg/pagination"
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

		projectResult, err := c.projectService.FindProjectById(r.Context(), &dto.FindProjectByIdParams{ID: projectId})
		if err != nil {
			if errors.Is(err, apperrors.ErrNotFound) {
				httputil.Render(w, r, apperrors.ErrHTTPNotFoundError())
				return
			}

			httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
			return
		}

		ctx := context.WithValue(r.Context(), ProjectCtxKey, projectResult.Project)
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
//	@param		limit	query		int64	false	"Amount of results to return"
//	@param		offset	query		int64	false	"Offset of results to return"
//	@success	200		{object}	dto.ListProjectsResponse
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/projects [get]
func (c *ProjectController) FindAllProjects(w http.ResponseWriter, r *http.Request) {
	limit, err := pagination.Limit(r)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	offset, err := pagination.Offset(r)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	projectsResult, err := c.projectService.FindAllProjects(r.Context(), &dto.FindAllProjectsParams{
		PaginationParams: &dto.PaginationParams{Limit: limit, Offset: offset},
	})
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToListProjectsResponse(projectsResult))
}

// GetProject godoc
//
//	@summary	Get a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		projectId	path		uint	true	"Project identifier"
//	@success	200			{object}	dto.ProjectResponse
//	@failure	400			{object}	apperrors.ErrResponse
//	@failure	404			{object}	apperrors.ErrResponse
//	@failure	500			{object}	apperrors.ErrResponse
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
//	@param		request	body		dto.CreateProjectRequest	true	"Create a project"
//	@success	201		{object}	dto.ProjectResponse
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/projects [post]
func (c *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	input := &dto.CreateProjectRequest{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	projectResult, err := c.projectService.CreateProject(r.Context(), &dto.CreateProjectParams{
		Name: input.Name,
		Slug: input.Slug,
	})
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	httputil.Render(w, r, dto.ToProjectResponse(projectResult.Project))
}

// UpdateProject godoc
//
//	@summary	Update a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		projectId	path		uint						true	"Project identifier"
//	@param		request		body		dto.UpdateProjectRequest	true	"Update a project"
//	@success	200			{object}	dto.ProjectResponse
//	@failure	400			{object}	apperrors.ErrResponse
//	@failure	404			{object}	apperrors.ErrResponse
//	@failure	500			{object}	apperrors.ErrResponse
//	@router		/api/v1/projects/{projectId} [put]
func (c *ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	project := getProject(r.Context())

	input := &dto.UpdateProjectRequest{
		Slug: project.Slug,
		Name: project.Name,
	}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	projectResult, err := c.projectService.UpdateProject(r.Context(), &dto.UpdateProjectParams{
		Slug: input.Slug,
		Name: input.Name,
		ID:   project.ID,
	})
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToProjectResponse(projectResult.Project))
}

// DeleteProject godoc
//
//	@summary	Delete a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		projectId	path	uint	true	"Project identifier"
//	@success	204
//	@failure	404	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/projects/{projectId} [delete]
func (c *ProjectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	project := getProject(r.Context())

	err := c.projectService.DeleteProject(r.Context(), &dto.DeleteProjectParams{ID: project.ID})
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}
