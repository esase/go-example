package routes

import (
	"net/http"

	"git.crg.one/scm/go/supplier-hub.git/internal/common/interfaces"
	"git.crg.one/scm/go/supplier-hub.git/internal/common/schema"
	"git.crg.one/scm/go/supplier-hub.git/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(router *gin.Engine) {
	cancelRouteHandlers := interfaces.CancelRouteHandlersMap{}

	// register platform handlers
	handlers.RegisterRouteHandlers(cancelRouteHandlers)

	// platforms routes
	router.POST("/:platform/cancel", func(c *gin.Context) {
		platformName := c.Param("platform")

		// check if there is a handler for received platform
		routeHandler, ok := cancelRouteHandlers[platformName]

		if !ok {
			c.AbortWithStatus(http.StatusNotImplemented)

			return
		}

		var request schema.CancelRequestParams

		// hydrate params
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid request body")

			return
		}

		routeHandler(c, request)
	})
}
