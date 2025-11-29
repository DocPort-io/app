package app

import (
	"app/pkg/controller"
	"app/pkg/service"
	"app/pkg/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB, fileStorage storage.FileStorage) http.Handler {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	fileService := service.NewFileService(db, fileStorage)
	projectService := service.NewProjectService(db)
	versionService := service.NewVersionService(db, fileService)

	projectController := controller.NewProjectController(projectService)
	versionController := controller.NewVersionController(versionService)
	fileController := controller.NewFileController(fileService)

	registerRoutes(router, projectController, versionController, fileController)

	return router
}
