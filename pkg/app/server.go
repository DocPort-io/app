package app

import (
	"app/pkg/database"
	"app/pkg/file"
	"app/pkg/project"
	"app/pkg/storage"
	"app/pkg/version"
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

	projectService := project.NewProjectService(queries)
	versionService := version.NewVersionService(queries)
	fileService := file.NewFileService(queries, fileStorage)

	projectController := project.NewProjectController(projectService)
	versionController := version.NewVersionController(versionService)
	fileController := file.NewFileController(fileService)

	registerRoutes(router, projectController, versionController, fileController)

	return router
}
