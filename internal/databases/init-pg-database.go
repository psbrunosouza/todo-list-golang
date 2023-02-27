package databases

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDB store the database instance
var PostgresDB *gorm.DB

// InitPGDatabase initialize postgres database
func InitPGDatabase(models ...interface{}) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error during database initialization")
	}

	if autoMigrateError := db.AutoMigrate(models...); autoMigrateError != nil {
		log.Fatalln("Error during auto migration")
	}

	PostgresDB = db
}
