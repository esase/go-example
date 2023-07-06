package routes

import (
	routes "git.crg.one/scm/go/supplier-hub.git/internal/suppliers"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(router *gin.Engine) {
	// supplier routes
	routes.InstallRoutes(router)
}
