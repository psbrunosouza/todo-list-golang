package ratings

import (
	"todo-list/internal/databases"

	"github.com/labstack/echo/v4"
)

func RatingRoutes(e *echo.Group) {
	ratingRepository := NewRatingRepository(databases.PostgresDB)
	ratingService := NewRatingService(ratingRepository)
	ratingHandler := NewRatingHandler(ratingService)

	e.GET("/ratings", ratingHandler.List)
	e.POST("/ratings", ratingHandler.Create)
	e.PUT("/ratings/:id", ratingHandler.Update)
	e.DELETE("/ratings/:id", ratingHandler.Delete)
	e.GET("/ratings/:id", ratingHandler.Find)
}
