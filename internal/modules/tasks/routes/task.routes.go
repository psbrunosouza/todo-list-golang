package task_routes

import (
	task_handlers "todo-list/internal/modules/tasks/handlers"

	"github.com/labstack/echo/v4"
)

func TasksRoutes(e *echo.Echo) {
	e.GET("/tasks", task_handlers.ListTaskHandler)
	e.POST("/tasks", task_handlers.CreateTaskHandler)
	e.PUT("/tasks/:id", task_handlers.UpdateTaskHandler)
	e.DELETE("/tasks/:id", task_handlers.DeleteTaskHandler)
	e.GET("/tasks/:id", task_handlers.FindTaskHandler)
}
