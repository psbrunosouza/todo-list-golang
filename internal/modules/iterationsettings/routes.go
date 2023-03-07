package iterationsettings

import "github.com/labstack/echo/v4"

func IterationSettingsRoutes(e *echo.Echo) {
	e.GET("/iteration-settings", ListIterationSettingsHandler)
	e.POST("/iteration-settings", CreateIterationSettingHandler)
	e.PUT("/iteration-settings/:id", UpdateIterationSettingHandler)
	e.DELETE("/iteration-settings/:id", DeleteIterationSettingHandler)
	e.GET("/iteration-settings/:id", FindIterationSettingHandler)
}
