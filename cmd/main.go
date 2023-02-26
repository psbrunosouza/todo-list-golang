package main

import (
	"log"
	"os"
	"todo-list/internal/databases"
	task_routes "todo-list/internal/modules/tasks/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if dotEnvError := godotenv.Load(); dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitPGDatabase()
}

func main() {
	echo := echo.New()

	task_routes.TasksRoutes(echo)

	echo.Logger.Fatal(echo.Start(os.Getenv("PORT")))
}
