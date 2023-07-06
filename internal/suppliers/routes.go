package suppliers

import (
	"git.crg.one/scm/go/supplier-hub.git/internal/suppliers/hertz"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(router *gin.Engine) {
	hertz.InstallRoutes(router)
}
