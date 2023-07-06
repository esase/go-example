package handlers

import (
	"git.crg.one/scm/go/supplier-hub.git/internal/common/interfaces"
	"git.crg.one/scm/go/supplier-hub.git/internal/handlers/foco"
	"git.crg.one/scm/go/supplier-hub.git/internal/handlers/hertz"
)

func RegisterRouteHandlers(cancelRouteHandlersMap interfaces.CancelRouteHandlersMap) {
	foco.RegisterRouteHandlers(cancelRouteHandlersMap)
	hertz.RegisterRouteHandlers(cancelRouteHandlersMap)
}
