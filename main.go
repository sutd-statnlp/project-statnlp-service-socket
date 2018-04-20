package main

import (
	"github.com/gin-gonic/gin"

	"./controller"
	"./handler"
)

func main() {
	router := setupRoutes()

	router.Run(":8220")
}

func setupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Next()
	})

	controller.InitHomeRoutes(router)
	handler.ConfigRequest(router)

	return router
}
