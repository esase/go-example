package hertz

import (
	"git.crg.one/scm/go/supplier-hub.git/internal/common/interfaces"
)

const PlatformName = "hertz"

func RegisterRouteHandlers(cancelRouteHandlersMap interfaces.CancelRouteHandlersMap) {
	cancelRouteHandlersMap[PlatformName] = CancelHandler
}
