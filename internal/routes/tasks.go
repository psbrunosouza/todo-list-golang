package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func createTask(context echo.Context) error {
	return context.JSON(http.StatusCreated, "")
}

func getTasks(context echo.Context) error {
	return context.JSON(http.StatusOK, "")
}

func updateTask(context echo.Context) error {
	return context.JSON(http.StatusOK, "")
}

func deleteTask(context echo.Context) error {
	return context.JSON(http.StatusOK, "")
}

func findTask(context echo.Context) error {
	return context.JSON(http.StatusOK, "")
}

func TasksRoute(e *echo.Echo) {
	e.GET("/tasks", getTasks)
	e.POST("/tasks", createTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)
	e.GET("/tasks/:id", findTask)
}
