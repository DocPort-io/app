package controller

import (
	"app/pkg/apperrors"
	"app/pkg/dto"
	"app/pkg/httputil"
	"app/pkg/service"
	"net/http"

	"github.com/go-chi/render"
)

type ProjectController struct {
	projectService *service.ProjectService
}

func NewProjectController(projectService *service.ProjectService) *ProjectController {
	return &ProjectController{projectService: projectService}
}

// FindAllProjects godoc
//
//	@summary	Find all projects
//	@tags		projects
//	@accept		json
//	@produce	json
//	@success	200	{object}	dto.ListProjectsResponseDto
//	@router		/projects [get]
func (c *ProjectController) FindAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, total, err := c.projectService.FindAllProjects(r.Context())
	if err != nil {
		render.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, dto.ToListProjectsResponse(projects, total))
}

// GetProject godoc
//
//	@summary	Get a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		id	path		uint	true	"Project identifier"
//	@success	200	{object}	dto.ProjectResponseDto
//	@router		/projects/{id} [get]
func (c *ProjectController) GetProject(w http.ResponseWriter, r *http.Request) {
	projectId, err := httputil.URLParamInt64(r, "projectId")
	if err != nil {
		render.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	project, err := c.projectService.FindProjectById(r.Context(), projectId)
	if err != nil {
		render.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, dto.ToProjectResponse(project))
}

// CreateProject godoc
//
//	@summary	Create a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		request	body		dto.CreateProjectDto	true	"Create a project"
//	@success	201		{object}	dto.ProjectResponseDto
//	@router		/projects [post]
func (c *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	input := &dto.CreateProjectDto{}
	if err := render.Bind(r, input); err != nil {
		render.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	project, err := c.projectService.CreateProject(r.Context(), input)
	if err != nil {
		render.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, dto.ToProjectResponse(project))
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
//	@router		/projects/{id} [put]
func (c *ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	input := &dto.UpdateProjectDto{}
	if err := render.Bind(r, input); err != nil {
		render.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	projectId, err := httputil.URLParamInt64(r, "projectId")
	if err != nil {
		render.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	project, err := c.projectService.UpdateProject(r.Context(), projectId, input)
	if err != nil {
		render.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, dto.ToProjectResponse(project))
}

// DeleteProject godoc
//
//	@summary	Delete a project
//	@tags		projects
//	@accept		json
//	@produce	json
//	@param		id	path	uint	true	"Project identifier"
//	@success	204
//	@router		/projects/{id} [delete]
func (c *ProjectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	projectId, err := httputil.URLParamInt64(r, "projectId")
	if err != nil {
		render.Render(w, r, apperrors.ErrBadRequestError(err))
		return
	}

	err = c.projectService.DeleteProject(r.Context(), projectId)
	if err != nil {
		render.Render(w, r, apperrors.ErrInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}
