package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")

	{
		api.GET("/init", func(ctx *gin.Context) {
			ctx.JSON(200, "fala zeze")
		})
	}

	router.Run(":8080")
}
