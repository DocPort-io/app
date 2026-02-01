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

	projectRepository := project.NewRepository(queries)
	versionRepository := version.NewRepository(queries)
	fileRepository := file.NewRepository(queries)

	projectService := project.NewService(projectRepository)
	versionService := version.NewVersionService(versionRepository)
	fileService := file.NewFileService(fileRepository, fileStorage)

	projectController := project.NewHandler(projectService)
	versionController := version.NewHandler(versionService)
	fileController := file.NewHandler(fileService)

	registerRoutes(router, projectController, versionController, fileController)

	return router
}
