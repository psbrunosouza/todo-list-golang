package tasks

import (
	"github.com/labstack/echo/v4"
)

func TasksRoutes(e *echo.Echo) {
	e.GET("/tasks", ListTaskHandler)
	e.POST("/tasks", CreateTaskHandler)
	e.PUT("/tasks/:id", UpdateTaskHandler)
	e.DELETE("/tasks/:id", DeleteTaskHandler)
	e.GET("/tasks/:id", FindTaskHandler)
}
