package example

import (
	"git.crg.one/scm/go/supplier-hub.git/internal/suppliers/example/handlers"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(router *gin.Engine) {
	group := router.Group("/example")
	{
		group.GET("/albums", handlers.ListHandler)
		group.POST("/albums", handlers.CreateHandler)
	}
}
