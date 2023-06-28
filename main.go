package main

import (
	_ "git.crg.one/scm/go/supplier-hub.git/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"git.crg.one/scm/go/supplier-hub.git/src/suppliers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	suppliers.InstallRoutes(router)
	router.Run()

	router.Run("localhost:8080")
}
