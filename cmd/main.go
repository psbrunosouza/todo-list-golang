package main

import (
	"log"
	"os"
	"todo-list/internal/databases"
	"todo-list/internal/entities"
	"todo-list/internal/modules/colors"
	"todo-list/internal/modules/iterations"
	"todo-list/internal/modules/iterationsettings"
	"todo-list/internal/modules/ratings"
	"todo-list/internal/modules/subtasks"
	"todo-list/internal/modules/tasks"
	"todo-list/internal/modules/users"
	"todo-list/internal/modules/workspaces"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if dotEnvError := godotenv.Load(); dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitPGDatabase(
		&entities.Task{},
		&entities.SubTask{},
		&entities.Workspace{},
		&entities.User{},
		&entities.Color{},
		&entities.Rating{},
		&entities.Iteration{},
		&entities.IterationSetting{},
	)
}

func main() {
	echo := echo.New()

	api := echo.Group("/v1")

	tasks.TasksRoutes(api)
	subtasks.SubTaskRoutes(api)
	workspaces.WorspaceRoutes(api)
	users.UserRoutes(api)
	ratings.RatingRoutes(api)
	colors.ColorRoutes(api)
	iterationsettings.IterationSettingsRoutes(api)
	iterations.IterationRoutes(api)

	echo.Logger.Fatal(echo.Start(os.Getenv("PORT")))
}
