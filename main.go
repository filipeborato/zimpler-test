package main

import (
	"zimpler-test/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/candy-zimpler", func(ctx *gin.Context) {
		ctx.JSON(200, controller.CandyStore(ctx))
	})
	server.Run(":8080")
}
