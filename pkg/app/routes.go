package app

import (
	"app/pkg/controller"
	_ "app/pkg/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title		DocPort.io API
//	@version	0.0.1

// @host		localhost:8080
// @basepath	/api/v1
func registerRoutes(router *gin.Engine, projectController *controller.ProjectController) {
	router.GET("/api/v1/projects", projectController.FindAllProjects)
	router.GET("/api/v1/projects/:id", projectController.GetProject)
	router.POST("/api/v1/projects", projectController.CreateProject)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
