package databases

import (
	"log"
	"os"
	tasks "todo-list/internal/modules/tasks/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func InitializePostgresDatabase() {
	database, databaseError := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if databaseError != nil {
		log.Fatalln("Error during database initialization")
	}

	if autoMigrateError := database.AutoMigrate(&tasks.Task{}); autoMigrateError != nil {
		log.Fatalln("Error during auto migration")
	}

	PostgresDB = database
}
