package main

import (
	"fmt"
	"os"

	routes "git.crg.one/scm/go/supplier-hub.git/internal"
	"git.crg.one/scm/go/supplier-hub.git/internal/middleware/openapi"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("configs/.env")

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := setupRouter()

	router.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(openapi.OpenapiValidator())

	routes.InstallRoutes(router)

	return router
}
