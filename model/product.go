package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id         int
	Name       string
	Price      float32
	Quantity   int
	Descripton string
}
