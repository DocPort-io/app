package app

import (
	"app/pkg/controller"
	"app/pkg/database"
	"app/pkg/service"
	"app/pkg/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func NewServer(queries *database.Queries, fileStorage storage.FileStorage) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/heartbeat"))
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	projectService := service.NewProjectService(queries)
	versionService := service.NewVersionService(queries)
	fileService := service.NewFileService(queries, fileStorage)

	projectController := controller.NewProjectController(projectService)
	versionController := controller.NewVersionController(versionService)
	fileController := controller.NewFileController(fileService)

	registerRoutes(router, projectController, versionController, fileController)

	return router
}
