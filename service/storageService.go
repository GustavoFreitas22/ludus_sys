package service

import (
	"github.com/GustavoFreitas22/ludus_sys/model"
	"github.com/GustavoFreitas22/ludus_sys/repository"
)

func ValidateIfItemExistInStorage(newItems []model.Product, id int) model.Storage {
	var validItems []model.Product
	storageItems := repository.ListStorageItems(id)

	for _, updatedItem := range newItems {
		for _, item := range storageItems {
			if item.Id != updatedItem.Id && item.Name != updatedItem.Name {
				validItems = append(validItems, updatedItem)
			}
		}
	}

	storageItems = append(storageItems, validItems...)

	return model.Storage{Items: storageItems}
}
