package app

import (
	"app/pkg/controller"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine, projectController *controller.ProjectController) {
	router.GET("/api/v1/projects", projectController.FindAllProjects)
	router.POST("/api/v1/projects", projectController.CreateProject)
}
