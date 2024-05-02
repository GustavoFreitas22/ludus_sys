package model

import (
	"time"

	"gorm.io/gorm"
)

type Sell struct {
	gorm.Model
	Date       time.Time
	StorageId  int
	Items      []Product
	TotalValue float32
}

type Historic struct {
	gorm.Model
	SellList []Sell
}
