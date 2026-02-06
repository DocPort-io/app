package app

import (
	"app/pkg/file"
	"app/pkg/platform/config"
	platformMiddleware "app/pkg/platform/middleware"
	"app/pkg/project"
	"app/pkg/user"
	"app/pkg/version"
	"log"
	"net"
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

// @securitydefinitions.oauth2.application OAuth2ClientCredentials
// @tokenUrl https://keycloak.docport.io/realms/docport-dev/protocol/openid-connect/token

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @authorizationurl https://keycloak.docport.io/realms/docport-dev/protocol/openid-connect/auth
// @tokenUrl https://keycloak.docport.io/realms/docport-dev/protocol/openid-connect/token
func NewServer() http.Server {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	fileStorage := NewFileStorage(cfg)
	queries := NewDatabase(cfg)

	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/heartbeat"))
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	authMiddleware, err := platformMiddleware.NewAuthMiddleware(cfg)
	if err != nil {
		log.Fatalf("creating auth middleware failed: %v", err)
	}

	projectRepository := project.NewRepository(queries)
	versionRepository := version.NewRepository(queries)
	fileRepository := file.NewRepository(queries)

	projectService := project.NewService(projectRepository)
	versionService := version.NewVersionService(versionRepository)
	fileService := file.NewFileService(fileRepository, fileStorage)

	projectHandler := project.NewHandler(projectService)
	versionHandler := version.NewHandler(versionService)
	fileHandler := file.NewHandler(fileService)
	userHandler := user.NewHandler(authMiddleware)

	router.Route("/api/v1", func(r chi.Router) {
		projectHandler.RegisterRoutes(r)
		versionHandler.RegisterRoutes(r)
		fileHandler.RegisterRoutes(r)
		userHandler.RegisterRoutes(r)
	})

	docs.SwaggerInfo.Host = viper.GetString("server.host")

	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	return http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Bind, cfg.Server.Port),
		Handler: router,
	}
}
