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
func registerRoutes(router *gin.Engine, projectController *controller.ProjectController, versionController *controller.VersionController) {
	router.GET("/api/v1/projects", projectController.FindAllProjects)
	router.GET("/api/v1/projects/:id", projectController.GetProject)
	router.POST("/api/v1/projects", projectController.CreateProject)
	router.PUT("/api/v1/projects/:id", projectController.UpdateProject)
	router.DELETE("/api/v1/projects/:id", projectController.DeleteProject)

	router.GET("/api/v1/versions", versionController.FindAllVersions)
	router.GET("/api/v1/versions/:id", versionController.GetVersion)
	router.POST("/api/v1/versions", versionController.CreateVersion)
	router.PUT("/api/v1/versions/:id", versionController.UpdateVersion)
	router.DELETE("/api/v1/versions/:id", versionController.DeleteVersion)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
