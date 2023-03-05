package users

import "github.com/labstack/echo/v4"

func UserRoutes(e *echo.Echo) {
	e.GET("/users", ListUsersHandler)
	e.POST("/users", CreateUsersHandler)
	e.PUT("/users/:id", UpdateUserHandler)
	e.DELETE("/users/:id", DeleteUserHandler)
	e.GET("/users/:id", FindUserHandler)
}
