package app

import (
	"app/pkg/file"
	"app/pkg/project"
	"app/pkg/version"

	"app/pkg/docs"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title		DocPort.io API
//	@version	0.0.1

// @host		localhost:8080
// @basepath	/
func registerRoutes(router *chi.Mux, projectController *project.Handler, versionController *version.Handler, fileController *file.Handler) {
	router.Route("/api/v1", func(r chi.Router) {
		projectController.RegisterRoutes(r)
		versionController.RegisterRoutes(r)
		fileController.RegisterRoutes(r)
	})

	docs.SwaggerInfo.Host = viper.GetString("server.host")

	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
}
