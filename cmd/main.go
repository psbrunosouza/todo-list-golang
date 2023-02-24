package main

import (
	"todo-list/internal/config"
	"todo-list/internal/databases"
	tasks "todo-list/internal/modules/tasks/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.InitializeEnv()
	databases.InitializePostgresDatabase()
	tasks.TasksRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
