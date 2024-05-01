package repository

import (
	"log"

	"github.com/GustavoFreitas22/ludus_sys/config"
	"github.com/GustavoFreitas22/ludus_sys/model"
)

func CreateNewProduct(product model.Product) {
	config.Datasource.Create(&product)
}

func FindProductByName(name string) model.Product {

	var product model.Product

	config.Datasource.Find(&product, "name = ?", name)

	log.Println("Find product: ", product.Name)

	return product
}

func FindAllProducts() []model.Product {
	var product []model.Product

	config.Datasource.Order("name").Find(&product)

	return product
}

func FindProductById(id int) model.Product {

	var product model.Product

	config.Datasource.Find(&product, "ID = ?", id)

	log.Println("Find product: ", product.ID)

	return product
}

func UpdateProduct(id int, dataUpdated model.Product) {
	var product model.Product

	config.Datasource.Where("ID = ?", id).First(&product)

	product.Descripton = dataUpdated.Descripton
	product.Name = dataUpdated.Name
	product.Price = dataUpdated.Price

	config.Datasource.Save(&product)
}

func DeleteProduct(id int) {
	config.Datasource.Delete(&model.Product{}, id)
}
