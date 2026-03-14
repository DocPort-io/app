package app

import (
	"app/pkg/api"
	"app/pkg/file"
	"app/pkg/platform/config"
	"app/pkg/platform/handler"
	platformMiddleware "app/pkg/platform/middleware"
	"app/pkg/platform/swagger"
	"app/pkg/project"
	"app/pkg/user"
	"app/pkg/version"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func NewServer() *http.Server {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	openapi, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("failed to get swagger spec: %v", err)
	}

	fileStorage := NewFileStorage(cfg)
	queries := NewDatabase(cfg.Database.DSN)

	projectRepository := project.NewRepository(queries)
	versionRepository := version.NewRepository(queries)
	fileRepository := file.NewRepository(queries)
	userRepository := user.NewRepository(queries)

	projectService := project.NewService(projectRepository)
	versionService := version.NewVersionService(versionRepository)
	fileService := file.NewFileService(fileRepository, fileStorage)
	userService := user.NewService(userRepository)

	authMiddleware, err := platformMiddleware.NewAuthMiddleware(cfg)
	if err != nil {
		log.Fatalf("creating auth middleware failed: %v", err)
	}

	oapiMiddleware := nethttpmiddleware.OapiRequestValidatorWithOptions(openapi, &nethttpmiddleware.Options{
		DoNotValidateServers: true,
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			handler.WriteError(w, statusCode, message)
		},
	})

	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/heartbeat"))
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	projectHandler := project.NewHandler(projectService)
	versionHandler := version.NewHandler(versionService)
	fileHandler := file.NewHandler(fileService)
	userHandler := user.NewHandler(userService, authMiddleware)

	router.Route("/api", func(r chi.Router) {
		r.Use(oapiMiddleware)
		projectHandler.RegisterRoutes(r)
		versionHandler.RegisterRoutes(r)
		fileHandler.RegisterRoutes(r)
		userHandler.RegisterRoutes(r)
	})

	swagger.SetupRoutes(router, openapi)

	return &http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Bind, cfg.Server.Port),
		Handler: router,
	}
}
