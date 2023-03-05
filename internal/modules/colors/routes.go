package colors

import "github.com/labstack/echo/v4"

func ColorRoutes(e *echo.Echo) {
	e.GET("/colors", ListColorsHandler)
	e.POST("/colors", CreateColorHandler)
	e.PUT("/colors/:id", UpdateColorHandler)
	e.DELETE("/colors/:id", DeleteColorHandler)
	e.GET("/colors/:id", FindColorHandler)
}
