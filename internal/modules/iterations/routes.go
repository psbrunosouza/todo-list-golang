package iterations

import "github.com/labstack/echo/v4"

func IterationRoutes(e *echo.Echo) {
	e.GET("/iterations", ListIterationsHandler)
	e.POST("/iterations", CreateIterationHandler)
	e.PUT("/iterations/:id", UpdateIterationHandler)
	e.DELETE("/iterations/:id", DeleteIterationHandler)
	e.GET("/iterations/:id", FindIterationHandler)
}
