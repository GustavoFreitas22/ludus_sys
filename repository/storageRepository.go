package repository

import (
	"github.com/GustavoFreitas22/ludus_sys/config"
	"github.com/GustavoFreitas22/ludus_sys/model"
)

func ListStorageItems(id int) []model.Product {
	storage := model.Storage{}
	config.Datasource.Order("items").Where("id = ?", id).Find(&storage)

	return storage.Items
}

func CreateStorage(storage model.Storage) {
	config.Datasource.Create(&storage)
}

func FindStorageByName(name string) model.Storage {
	storage := model.Storage{}

	config.Datasource.Where("name = ?", name).Find(&storage)

	return storage
}

func DeleteStorage(id int) {
	storage := model.Storage{}
	config.Datasource.Where("id = ?", id).Delete(&storage)
}

func UpdateStorageItems(storageUpdated model.Storage, id int) {
	var storage model.Storage

	config.Datasource.Where("id = ?", id).First(&storage)

	storage.Items = storageUpdated.Items

	config.Datasource.Save(&storage)
}
