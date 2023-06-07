package service

import (
	"log"

	"github.com/GustavoFreitas22/ludus_sys/repository"
)

func ValidateNewModality(name string) bool {

	exist := validateIfExist(name)

	if !exist {
		return false
	}

	repository.CreateNewModality(name)
	return true
}

func ValidateEditModality(name string) bool {
	exist := validateIfExist(name)
	return exist
}

func validateIfExist(name string) bool {
	modalities := repository.FindAllModalities()

	for _, modality := range modalities {
		if modality.Name == name {
			log.Println("Invalid: Modality already exists")
			return false
		}
	}
	return true
}
