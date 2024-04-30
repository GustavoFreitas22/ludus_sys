package controller

import (
	"fmt"
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

		api.GET("/product/all", func(ctx *gin.Context) {
			allProducts := repository.FindAllProducts()

			ctx.JSON(200, allProducts)
		})

		api.GET("/product/:id", func(ctx *gin.Context) {

			param := ctx.Param("id")

			id, err := strconv.Atoi(param)
			if err != nil {
				ctx.JSON(402, "BAD REQUEST")
				return
			}

			product := repository.FindProductById(id)

			ctx.JSON(200, product.Name)
		})

		api.GET("/product/name/:name", func(ctx *gin.Context) {

			name := ctx.Param("name")

			fmt.Println(name)

			product := repository.FindProductByName(name)

			ctx.JSON(200, product)
		})

		api.POST("/product/", func(ctx *gin.Context) {
			var newProduct model.Product

			ctx.BindJSON(&newProduct)

			validName := service.ValidateIfExist(newProduct.Name)

			if !validName {
				ctx.JSON(400, "not created, product aready exist")
				return
			}

			ctx.JSON(201, "created with success")
		})

		api.PUT("/product/:id/:name", func(ctx *gin.Context) {
			idParam := ctx.Param("id")

			var updatedDataProduct model.Product
			ctx.BindJSON(&updatedDataProduct)

			id, err := strconv.Atoi(idParam)
			if err != nil {
				ctx.JSON(402, "Invalid id")
				return
			}

			validName := service.ValidateIfExist(updatedDataProduct.Name)

			if !validName {
				ctx.JSON(400, "Product with this name already exist in database")
				return
			}

			repository.UpdateProduct(id, updatedDataProduct)

			ctx.JSON(200, "Edited")
		})

		api.DELETE("/product/:id", func(ctx *gin.Context) {
			idParam := ctx.Param("id")

			id, err := strconv.Atoi(idParam)
			if err != nil {
				ctx.JSON(402, "Invalid Id")
				return
			}

			repository.DeleteProduct(id)

			ctx.JSON(200, "Deleted")
		})

	}

	router.Run(":8080")
}
