package example

import (
	"git.crg.one/scm/go/supplier-hub.git/src/suppliers/example/handlers/create"
	"git.crg.one/scm/go/supplier-hub.git/src/suppliers/example/handlers/list"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(router *gin.Engine) {
	group := router.Group("/example")
	{
		group.GET("/albums", list.Handler)
		group.POST("/albums", create.Handler)
	}
}
