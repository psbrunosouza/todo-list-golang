package databases

import (
	"log"
	"os"
	tasks "todo-list/internal/modules/tasks/task_models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func InitializePostgresDatabase() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error during database initialization")
	}

	if autoMigrateError := db.AutoMigrate(&tasks.Task{}); autoMigrateError != nil {
		log.Fatalln("Error during auto migration")
	}

	PostgresDB = db
}
