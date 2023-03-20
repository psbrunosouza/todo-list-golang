package colors

import (
	"todo-list/internal/databases"

	"github.com/labstack/echo/v4"
)

func ColorRoutes(e *echo.Group) {
	colorRepository := NewColorRepository(databases.PostgresDB)
	colorService := NewColorService(colorRepository)
	colorHandler := NewColorHandler(colorService)

	e.GET("/colors", colorHandler.List)
	e.POST("/colors", colorHandler.Create)
	e.PUT("/colors/:id", colorHandler.Update)
	e.DELETE("/colors/:id", colorHandler.Delete)
	e.GET("/colors/:id", colorHandler.Find)
}
