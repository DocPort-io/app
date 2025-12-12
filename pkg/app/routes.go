package app

import (
	"app/pkg/controller"

	_ "app/pkg/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title		DocPort.io API
//	@version	0.0.1

// @host		localhost:8080
// @basepath	/api/v1
func registerRoutes(router *chi.Mux, projectController *controller.ProjectController, versionController *controller.VersionController, fileController *controller.FileController) {
	router.Route("/api/v1", func(r chi.Router) {

		r.Route("/projects", func(r chi.Router) {
			r.Get("/", projectController.FindAllProjects)
			r.Post("/", projectController.CreateProject)

			r.Route("/{projectId}", func(r chi.Router) {
				r.Get("/", projectController.GetProject)
				r.Put("/", projectController.UpdateProject)
				r.Delete("/", projectController.DeleteProject)
			})
		})

		r.Route("/versions", func(r chi.Router) {
			r.Get("/", versionController.FindAllVersions)
			r.Post("/", versionController.CreateVersion)

			r.Route("/{versionId}", func(r chi.Router) {
				r.Get("/", versionController.GetVersion)
				r.Put("/", versionController.UpdateVersion)
				r.Delete("/", versionController.DeleteVersion)
				r.Post("/attach-file", versionController.AttachFileToVersion)
				r.Post("/detach-file", versionController.DetachFileFromVersion)
			})
		})

		r.Route("/files", func(r chi.Router) {
			r.Get("/", fileController.FindAllFiles)
			r.Post("/", fileController.CreateFile)

			r.Route("/{fileId}", func(r chi.Router) {
				r.Get("/", fileController.GetFile)
				r.Post("/upload", fileController.UploadFile)
				r.Delete("/", fileController.DeleteFile)
			})
		})
	})

	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
}
