package main

import (
	"fmt"
	"os"

	routes "git.crg.one/scm/go/supplier-hub.git/internal"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.InstallRoutes(router)

	return router
}

func main() {
	_ = godotenv.Load("configs/.env")

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := setupRouter()

	router.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
