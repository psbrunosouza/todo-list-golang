package ratings

import "github.com/labstack/echo/v4"

func RatingRoutes(e *echo.Echo) {
	e.GET("/ratings", ListRatingsHandler)
	e.POST("/ratings", CreateRatingHandler)
	e.PUT("/ratings/:id", UpdateRatingHandler)
	e.DELETE("/ratings/:id", DeleteRatingHandler)
	e.GET("/ratings/:id", FindRatingHandler)
}
