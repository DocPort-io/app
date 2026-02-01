package app

import (
	"app/pkg/database"
	"app/pkg/file"
	"app/pkg/project"
	"app/pkg/storage"
	"app/pkg/user"
	"app/pkg/version"
	"net/http"

	"app/pkg/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title		DocPort.io API
//	@version	0.0.1

// @host		localhost:8080
// @basepath	/
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

	projectHandler := project.NewHandler(projectService)
	versionHandler := version.NewHandler(versionService)
	fileHandler := file.NewHandler(fileService)
	userHandler := user.NewHandler()

	router.Route("/api/v1", func(r chi.Router) {
		projectHandler.RegisterRoutes(r)
		versionHandler.RegisterRoutes(r)
		fileHandler.RegisterRoutes(r)
		userHandler.RegisterRoutes(r)
	})

	docs.SwaggerInfo.Host = viper.GetString("server.host")

	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	return router
}
