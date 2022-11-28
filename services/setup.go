package services

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("guitars.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	err = database.AutoMigrate(&Guitar{})
	if err != nil {
		return
	}

	DB = database
}
