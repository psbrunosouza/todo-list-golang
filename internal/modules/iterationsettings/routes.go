package iterationsettings

import (
	"todo-list/internal/databases"

	"github.com/labstack/echo/v4"
)

func IterationSettingsRoutes(e *echo.Group) {
	iterationSettingRepository := NewIterationSettingRepository(databases.PostgresDB)
	iterationSettingService := NewIterationSettingService(iterationSettingRepository)
	iterationSettingHandler := NewIterationSettingHandler(iterationSettingService)

	e.GET("/iteration-settings", iterationSettingHandler.List)
	e.POST("/iteration-settings", iterationSettingHandler.Create)
	e.PUT("/iteration-settings/:id", iterationSettingHandler.Update)
	e.DELETE("/iteration-settings/:id", iterationSettingHandler.Delete)
	e.GET("/iteration-settings/:id", iterationSettingHandler.Find)
}
