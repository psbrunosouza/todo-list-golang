package databases

import (
	"log"
	"os"
	task_models "todo-list/internal/modules/tasks/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func InitializePostgresDatabase() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error during database initialization")
	}

	if autoMigrateError := db.AutoMigrate(&task_models.Task{}); autoMigrateError != nil {
		log.Fatalln("Error during auto migration")
	}

	PostgresDB = db
}
