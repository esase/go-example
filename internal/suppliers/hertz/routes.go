package hertz

import (
	"git.crg.one/scm/go/supplier-hub.git/internal/suppliers/hertz/handlers"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(router *gin.Engine) {
	group := router.Group("/hertz")
	{
		group.POST("/cancel", handlers.CancelHandler)
	}
}
