package foco

import (
	"git.crg.one/scm/go/supplier-hub.git/internal/common/interfaces"
)

const PlatformName = "foco"

func RegisterRouteHandlers(cancelRouteHandlersMap interfaces.CancelRouteHandlersMap) {
	cancelRouteHandlersMap[PlatformName] = CancelHandler
}
