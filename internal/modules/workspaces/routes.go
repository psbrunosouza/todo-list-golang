package workspaces

import "github.com/labstack/echo/v4"

func WorspaceRoutes(e *echo.Echo) {
	e.GET("/workspaces", ListWorkspacesHandler)
	e.POST("/workspaces", CreateWorkspaceHandler)
	e.PUT("/workspaces/:id", UpdateWorkspaceHandler)
	e.DELETE("/workspaces/:id", DeleteWorkspaceHandler)
	e.GET("/workspaces/:id", FindWorkspaceHandler)
}
