package app

import (
	"app/pkg/controller"
	"app/pkg/database"
	"app/pkg/service"
	"app/pkg/storage"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(db *sql.DB, queries *database.Queries, fileStorage storage.FileStorage) http.Handler {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	fileService := service.NewFileService(queries, fileStorage)
	projectService := service.NewProjectService(queries)
	versionService := service.NewVersionService(queries, fileService)

	projectController := controller.NewProjectController(projectService)
	versionController := controller.NewVersionController(versionService)
	fileController := controller.NewFileController(fileService)

	registerRoutes(router, projectController, versionController, fileController)

	return router
}
