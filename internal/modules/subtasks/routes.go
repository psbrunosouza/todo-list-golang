package subtasks

import (
	"github.com/labstack/echo/v4"
)

func SubTaskRoutes(e *echo.Echo) {
	e.GET("/sub-tasks", ListSubTasksHandler)
	e.POST("/sub-tasks", CreateSubTaskHandler)
	e.PUT("/sub-tasks/:id", UpdateSubTaskHandler)
	e.DELETE("/sub-tasks/:id", DeleteSubTaskHandler)
	e.GET("/sub-tasks/:id", FindSubTaskHandler)
}
