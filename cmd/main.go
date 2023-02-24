package main

import (
	"log"
	"os"
	"todo-list/internal/databases"
	taskRoutes "todo-list/internal/modules/tasks/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if dotEnvError := godotenv.Load(); dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitializePostgresDatabase()
}

func main() {
	echo := echo.New()

	taskRoutes.TasksRoutes(echo)

	echo.Logger.Fatal(echo.Start(os.Getenv("PORT")))
}
