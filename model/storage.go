package model

import (
	"gorm.io/gorm"
)

type Storage struct {
	gorm.Model
	Id    int
	Name  string
	Items []Product
}
