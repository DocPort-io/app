package controller

import (
	"app/pkg/dto"
	"app/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProjectController struct {
	db *gorm.DB
}

func NewProjectController(db *gorm.DB) *ProjectController {
	return &ProjectController{db: db}
}

// FindAllProjects godoc
//
//	@summary	Find all projects
//	@tags		projects
//	@accept		json
//	@produce	json
//	@success	200	{object}	dto.ListProjectsResponseDto
//	@router		/projects [get]
func (c *ProjectController) FindAllProjects(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	projects, err := gorm.G[model.Project](c.db).Find(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, dto.ToListProjectsResponse(projects, int64(len(projects))))
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
func (c *ProjectController) GetProject(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	id := ginCtx.Param("id")

	project, err := gorm.G[model.Project](c.db).Preload("Versions", nil).Where("id = ?", id).First(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, dto.ToProjectResponse(&project))
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
func (c *ProjectController) CreateProject(ginCtx *gin.Context) {
	var projectDto dto.CreateProjectDto
	if err := ginCtx.ShouldBindJSON(&projectDto); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := ginCtx.Request.Context()

	project := projectDto.ToModel()

	err := gorm.G[model.Project](c.db).Create(ctx, project)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusCreated, dto.ToProjectResponse(project))
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
func (c *ProjectController) UpdateProject(ginCtx *gin.Context) {
	var updateProjectDto dto.UpdateProjectDto
	if err := ginCtx.ShouldBindJSON(&updateProjectDto); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := ginCtx.Request.Context()

	id := ginCtx.Param("id")
	updateProject := updateProjectDto.ToModel()

	rowsAffected, err := gorm.G[model.Project](c.db).Where("id = ?", id).Updates(ctx, *updateProject)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		ginCtx.JSON(http.StatusNotFound, gin.H{
			"error": "project not found",
		})
		return
	}

	updatedProject, err := gorm.G[model.Project](c.db).Where("id = ?", id).First(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, dto.ToProjectResponse(&updatedProject))
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
func (c *ProjectController) DeleteProject(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	id := ginCtx.Param("id")

	rowsAffected, err := gorm.G[model.Project](c.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		ginCtx.JSON(http.StatusNotFound, gin.H{
			"error": "project not found",
		})
		return
	}

	ginCtx.JSON(http.StatusNoContent, nil)
}
