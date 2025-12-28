package config

import (
	"expense-tracker/internals/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("expense.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!", err)

	}
	err = database.AutoMigrate(&models.Expense{})
	if err != nil {
		log.Fatal("Migration failed", err)
	}
	DB = database
}
