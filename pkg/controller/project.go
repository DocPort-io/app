package controller

import (
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

func (c *ProjectController) FindAllProjects(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	projects, err := gorm.G[model.Project](c.db).Preload("Versions", nil).Find(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

type ProjectCreate struct {
	Slug string `json:"slug" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (c *ProjectController) CreateProject(ginCtx *gin.Context) {
	var project ProjectCreate
	if err := ginCtx.ShouldBindJSON(&project); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := ginCtx.Request.Context()

	err := gorm.G[model.Project](c.db).Create(ctx, &model.Project{
		Slug: project.Slug,
		Name: project.Name,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"message": "project created",
	})
}
