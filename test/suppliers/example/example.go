package example

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func RunAllTests(router *gin.Engine, t *testing.T) {
	testGettingAlbumsRoute(router, t)
	testCreatingAlbumRouteUsingEmptyBody(router, t)
	testCreatingAlbumRoute(router, t)
}

func testGettingAlbumsRoute(router *gin.Engine, t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/example/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `[
		{
			"id": "1",
			"title": "Blue Train",
			"artist": "John Coltrane",
			"price": 56.99
		},
		{
			"id": "2",
			"title": "Jeru",
			"artist": "Gerry Mulligan",
			"price": 17.99
		},
		{
			"id": "3",
			"title": "Sarah Vaughan and Clifford Brown",
			"artist": "Sarah Vaughan",
			"price": 39.99
		}
	]`, w.Body.String())
}

func testCreatingAlbumRouteUsingEmptyBody(router *gin.Engine, t *testing.T) {
	w := httptest.NewRecorder()

	body := []byte(`{}`)

	req, _ := http.NewRequest("POST", "/example/albums", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.JSONEq(t, `{
		"errors": [
			{
				"field": "Title",
				"message": "This field is required"
			},
			{
				"field": "Artist",
				"message": "This field is required"
			},
			{
				"field": "Price",
				"message": "This field is required"
			}
		]
	}`, w.Body.String())
}

func testCreatingAlbumRoute(router *gin.Engine, t *testing.T) {
	w := httptest.NewRecorder()

	body := []byte(`{
		"title" : "test",
		"artist": "www",
		"price": 100.22
	}`)

	req, _ := http.NewRequest("POST", "/example/albums", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.JSONEq(t, `{
		"id": "1",
		"title" : "test",
		"artist": "www",
		"price": 100.22
	}`, w.Body.String())
}
