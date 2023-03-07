package subtasks

import (
	"github.com/labstack/echo/v4"
)

func SubTaskRoutes(e *echo.Echo) {
	e.GET("/subtasks", ListSubTasksHandler)
	e.POST("/subtasks", CreateSubTaskHandler)
	e.PUT("/subtasks/:id", UpdateSubTaskHandler)
	e.DELETE("/subtasks/:id", DeleteSubTaskHandler)
	e.GET("/subtasks/:id", FindSubTaskHandler)
}
