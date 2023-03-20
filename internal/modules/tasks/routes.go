package tasks

import (
	"todo-list/internal/databases"

	"github.com/labstack/echo/v4"
)

func TasksRoutes(e *echo.Group) {

	taskRepository := NewTaskRepository(databases.PostgresDB)
	taskService := NewSubTaskService(taskRepository)
	taskHandler := NewTaskHandler(taskService)

	e.GET("/tasks", taskHandler.List)
	e.POST("/tasks", taskHandler.Create)
	e.PUT("/tasks/:id", taskHandler.Update)
	e.DELETE("/tasks/:id", taskHandler.Delete)
	e.GET("/tasks/:id", taskHandler.Find)
}
