package config

import (
	"github.com/GustavoFreitas22/ludus_sys/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Datasource *gorm.DB

func InitDatabase() {
	result, err := gorm.Open(sqlite.Open(".ludus.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = result.AutoMigrate(&model.Product{}, &model.Storage{})
	if err != nil {
		panic("failed to start database")
	}

	println("Database connect")

	Datasource = result
}
