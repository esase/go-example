package hertz

import (
	"fmt"
	"net/http"

	"git.crg.one/scm/go/supplier-hub.git/internal/common/schema"
	"github.com/gin-gonic/gin"
)

func CancelHandler(c *gin.Context, requestParams schema.CancelRequestParams) {
	fmt.Println("Hertz")
	fmt.Println(requestParams.Configuration.AsHertzConfiguration())

	var response schema.CancelResponse

	c.IndentedJSON(http.StatusCreated, response)
}
