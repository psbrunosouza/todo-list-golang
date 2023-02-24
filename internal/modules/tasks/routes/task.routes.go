package tasks

import (
	taskHandlers "todo-list/internal/modules/tasks/handlers"

	"github.com/labstack/echo/v4"
)

func TasksRoutes(e *echo.Echo) {
	// e.GET("/tasks", tasks.ListTaskHandler)
	e.POST("/tasks", taskHandlers.CreateTaskHandler)
	// e.PUT("/tasks/:id", tasks.UpdateTaskHandler)
	// e.DELETE("/tasks/:id", tasks.DeleteTaskHandler)
	// e.GET("/tasks/:id", tasks.FindTaskHandler)
}
