package app

import (
	"app/pkg/controller"
	"app/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) http.Handler {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	versionService := service.NewVersionService(db)

	projectController := controller.NewProjectController(db)
	versionController := controller.NewVersionController(versionService)

	registerRoutes(router, projectController, versionController)

	return router
}
