package users

import (
	"todo-list/internal/databases"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {

	userRepository := NewUserRepository(databases.PostgresDB)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	e.GET("/users", userHandler.List)
	e.POST("/users", userHandler.Create)
	e.PUT("/users/:id", userHandler.Update)
	e.DELETE("/users/:id", userHandler.Delete)
	e.GET("/users/:id", userHandler.Find)
}
