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
//	@param		id	path		uint	true	"Project ID"
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
