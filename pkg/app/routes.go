package app

import (
	"app/pkg/file"
	"app/pkg/paginate"
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
func registerRoutes(router *chi.Mux, projectController *project.Handler, versionController *version.VersionController, fileController *file.FileController) {
	router.Route("/api/v1", func(r chi.Router) {
		projectController.RegisterRoutes(r)

		r.Route("/versions", func(r chi.Router) {
			r.With(paginate.Paginate).Get("/", versionController.FindAllVersions)
			r.Post("/", versionController.CreateVersion)

			r.Route("/{versionId}", func(r chi.Router) {
				r.Use(versionController.VersionCtx)
				r.Get("/", versionController.GetVersion)
				r.Put("/", versionController.UpdateVersion)
				r.Delete("/", versionController.DeleteVersion)
				r.Post("/attach-file", versionController.AttachFileToVersion)
				r.Post("/detach-file", versionController.DetachFileFromVersion)
			})
		})

		r.Route("/files", func(r chi.Router) {
			r.With(paginate.Paginate).Get("/", fileController.FindAllFiles)
			r.Post("/", fileController.CreateFile)

			r.Route("/{fileId}", func(r chi.Router) {
				r.Use(fileController.FileCtx)
				r.Get("/", fileController.GetFile)
				r.Post("/upload", fileController.UploadFile)
				r.Get("/download", fileController.DownloadFile)
				r.Delete("/", fileController.DeleteFile)
			})
		})
	})

	docs.SwaggerInfo.Host = viper.GetString("server.host")

	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
}
