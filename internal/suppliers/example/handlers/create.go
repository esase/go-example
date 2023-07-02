package handlers

import (
	"net/http"

	"git.crg.one/scm/go/supplier-hub.git/internal/suppliers/example/models"
	"git.crg.one/scm/go/supplier-hub.git/internal/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateAlbum godoc
// @Summary                    Create Album
// @Description                Add a new Album
// @Tags                       album
// @Accept                     json
// @Produce                    json
// @Param                      album body     models.Request true "Album Data"
// @Success                    200  {object} models.Response
// @Router                     /example/albums [post]
func CreateHandler(c *gin.Context) {
	var request models.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.GetErrorList(err)})

		return
	}

	response := models.Response{
		ID:     "1",
		Title:  request.Title,
		Artist: request.Artist,
		Price:  request.Price,
	}

	c.IndentedJSON(http.StatusCreated, response)
}
