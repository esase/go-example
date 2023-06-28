package suppliers

import (
	"git.crg.one/scm/go/supplier-hub.git/src/suppliers/example"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(router *gin.Engine) {
	example.InstallRoutes(router)
}
