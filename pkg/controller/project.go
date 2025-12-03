package controller

import (
	"app/pkg/dto"
	"app/pkg/service"
	"app/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
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
func (c *ProjectController) FindAllProjects(ctx *gin.Context) {
	projects, err := c.projectService.FindAllProjects(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToListProjectsResponse(projects, int64(len(projects))))
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
func (c *ProjectController) GetProject(ctx *gin.Context) {
	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	project, err := c.projectService.FindProjectById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToProjectResponse(project))
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
func (c *ProjectController) CreateProject(ctx *gin.Context) {
	var input dto.CreateProjectDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	project, err := c.projectService.CreateProject(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToProjectResponse(project))
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
func (c *ProjectController) UpdateProject(ctx *gin.Context) {
	var input dto.UpdateProjectDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	project, err := c.projectService.UpdateProject(ctx.Request.Context(), id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToProjectResponse(project))
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
func (c *ProjectController) DeleteProject(ctx *gin.Context) {
	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	err = c.projectService.DeleteProject(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
