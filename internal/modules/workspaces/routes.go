package workspaces

import (
	"todo-list/internal/databases"

	"github.com/labstack/echo/v4"
)

func WorspaceRoutes(e *echo.Group) {
	workspaceRepository := NewWorkspaceRepository(databases.PostgresDB)
	workspaceService := NewWorkspaceService(workspaceRepository)
	workspaceHandler := NewWorkspaceHandler(workspaceService)

	e.GET("/workspaces", workspaceHandler.List)
	e.POST("/workspaces", workspaceHandler.Create)
	e.PUT("/workspaces/:id", workspaceHandler.Update)
	e.DELETE("/workspaces/:id", workspaceHandler.Delete)
	e.GET("/workspaces/:id", workspaceHandler.Find)
}
