package interfaces

import (
	"git.crg.one/scm/go/supplier-hub.git/internal/common/schema"
	"github.com/gin-gonic/gin"
)

type CancelRouteHandlersMap map[string]func(c *gin.Context, requestParams schema.CancelRequestParams)
