package app

import (
	"app/pkg/controller"
	"app/pkg/database"
	"app/pkg/service"
	"app/pkg/storage"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewServer(db *sql.DB, queries *database.Queries, fileStorage storage.FileStorage) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	projectService := service.NewProjectService(queries)
	versionService := service.NewVersionService(queries)
	fileService := service.NewFileService(queries, fileStorage)

	projectController := controller.NewProjectController(projectService)
	versionController := controller.NewVersionController(versionService)
	fileController := controller.NewFileController(fileService)

	registerRoutes(router, projectController, versionController, fileController)

	return router
}
