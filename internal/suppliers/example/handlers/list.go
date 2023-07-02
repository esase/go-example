package handlers

import (
	"net/http"

	"git.crg.one/scm/go/supplier-hub.git/internal/suppliers/example/models"
	"github.com/gin-gonic/gin"
)

// GetAllAlbums godoc
// @Summary                    Get Albums
// @Description                Get All Albums
// @Tags                       album
// @Produce                    json
// @Success                    200  {array} models.Response
// @Router                     /example/albums [get]
func ListHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []models.Response{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	})
}
