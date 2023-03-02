package main

import (
	"log"
	"os"
	"todo-list/internal/databases"
	"todo-list/internal/models"
	"todo-list/internal/modules/subtasks"
	"todo-list/internal/modules/tasks"
	"todo-list/internal/modules/workspaces"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if dotEnvError := godotenv.Load(); dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitPGDatabase(&models.Task{}, &models.SubTask{}, &models.Workspace{})
}

func main() {
	echo := echo.New()

	tasks.TasksRoutes(echo)
	subtasks.SubTaskRoutes(echo)
	workspaces.WorspaceRoutes(echo)

	echo.Logger.Fatal(echo.Start(os.Getenv("PORT")))
}
