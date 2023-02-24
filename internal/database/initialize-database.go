package database

import (
	"log"
	"os"
	"todo-list/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitializeDatabase() {
	database, databaseError := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if databaseError != nil {
		log.Fatalln("Error during database initialization")
	}

	if autoMigrateError := database.AutoMigrate(&models.Task{}); autoMigrateError != nil {
		log.Fatalln("Error during auto migration")
	}

	Database = database
}
