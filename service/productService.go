package service

import (
	"github.com/GustavoFreitas22/ludus_sys/model"
	"github.com/GustavoFreitas22/ludus_sys/repository"
)

func ValidateIfExist(name string) bool {
	product := repository.FindProductByName(name)
	return product != (model.Product{})
}
