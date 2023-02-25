package tasks

import (
	taskHandlers "todo-list/internal/modules/tasks/handlers"

	"github.com/labstack/echo/v4"
)

func TasksRoutes(e *echo.Echo) {
	e.GET("/tasks", taskHandlers.ListTaskHandler)
	e.POST("/tasks", taskHandlers.CreateTaskHandler)
	e.PUT("/tasks/:id", taskHandlers.UpdateTaskHandler)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTaskHandler)
	e.GET("/tasks/:id", taskHandlers.FindTaskHandler)
}
