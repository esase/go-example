package handlers

import (
	"fmt"
	"net/http"

	"git.crg.one/scm/go/supplier-hub.git/internal/suppliers/base/models"
	"github.com/gin-gonic/gin"
)

func CancelHandler(c *gin.Context) {
	var request models.CancelRequestParams

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid request body")

		return
	}

	fmt.Println(request)

	var response models.CancelResponse

	c.IndentedJSON(http.StatusCreated, response)
}
