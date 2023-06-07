package controller

import (
	"strconv"

	"github.com/GustavoFreitas22/ludus_sys/model"
	"github.com/GustavoFreitas22/ludus_sys/repository"
	"github.com/GustavoFreitas22/ludus_sys/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")

	{

		api.GET("/modality/all", func(ctx *gin.Context) {
			allModalities := repository.FindAllModalities()

			ctx.JSON(200, allModalities)
		})

		api.GET("/modality/:id", func(ctx *gin.Context) {

			param := ctx.Param("id")

			id, err := strconv.Atoi(param)
			if err != nil {
				ctx.JSON(402, "BAD REQUEST")
				return
			}

			modality := repository.FindModalityById(id)

			ctx.JSON(200, modality.Name)
		})

		api.POST("/modality/", func(ctx *gin.Context) {
			var newModality model.Modality

			ctx.BindJSON(&newModality)

			validName := service.ValidateNewModality(newModality.Name)

			if !validName {
				ctx.JSON(400, "not created")
				return
			}

			ctx.JSON(201, "created with success")
		})

		api.PUT("/modality/:id/:name", func(ctx *gin.Context) {
			idParam := ctx.Param("id")
			name := ctx.Param("name")

			id, err := strconv.Atoi(idParam)
			if err != nil {
				ctx.JSON(402, "BAD REQUEST")
				return
			}

			validName := service.ValidateEditModality(name)

			if !validName {
				ctx.JSON(400, "Name already exist in database")
				return
			}

			repository.UpdateModality(id, name)

			ctx.JSON(200, "Edited")
		})

		api.DELETE("/modality/:id", func(ctx *gin.Context) {
			idParam := ctx.Param("id")

			id, err := strconv.Atoi(idParam)
			if err != nil {
				ctx.JSON(402, "BAD REQUEST")
				return
			}

			repository.DeleteModality(id)

			ctx.JSON(200, "Deleted")
		})

	}

	{
		api.GET("/fala", func(ctx *gin.Context) {
			ctx.JSON(200, "testando e funcioanando")
		})
	}

	router.Run(":8080")
}
