package app

import (
	"app/pkg/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) http.Handler {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	projectController := controller.NewProjectController(db)

	registerRoutes(router, projectController)

	return router
}
