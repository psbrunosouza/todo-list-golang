package main

import (
	"todo-list/internal/config"
	"todo-list/internal/database"
	"todo-list/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.InitializeEnv()
	database.InitializeDatabase()
	routes.TasksRoute(e)

	e.Logger.Fatal(e.Start(":1323"))
}
