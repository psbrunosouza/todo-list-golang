package subtasks

import (
	"todo-list/internal/databases"

	"github.com/labstack/echo/v4"
)

func SubTaskRoutes(e *echo.Group) {
	subTaskRepository := NewSubTaskRepository(databases.PostgresDB)
	subTaskService := NewSubTaskService(subTaskRepository)
	subTaskHandler := NewSubTaskHandler(subTaskService)

	e.GET("/sub-tasks", subTaskHandler.List)
	e.POST("/sub-tasks", subTaskHandler.Create)
	e.PUT("/sub-tasks/:id", subTaskHandler.Update)
	e.DELETE("/sub-tasks/:id", subTaskHandler.Delete)
	e.GET("/sub-tasks/:id", subTaskHandler.Find)
}
