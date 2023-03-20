package iterations

import (
	"todo-list/internal/databases"

	"github.com/labstack/echo/v4"
)

func IterationRoutes(e *echo.Group) {
	iterationRepository := NewIterationRepository(databases.PostgresDB)
	iterationService := NewIterationService(iterationRepository)
	iterationHandler := NewIterationHandler(iterationService)

	e.GET("/iteration", iterationHandler.List)
	e.POST("/iteration", iterationHandler.Create)
	e.PUT("/iteration/:id", iterationHandler.Update)
	e.DELETE("/iteration/:id", iterationHandler.Delete)
	e.GET("/iteration/:id", iterationHandler.Find)
}
