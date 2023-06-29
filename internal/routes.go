package routes

import (
	_ "git.crg.one/scm/go/supplier-hub.git/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	routes "git.crg.one/scm/go/supplier-hub.git/internal/suppliers"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(router *gin.Engine) {
	// system routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// supplier routes
	routes.InstallRoutes(router)
}
