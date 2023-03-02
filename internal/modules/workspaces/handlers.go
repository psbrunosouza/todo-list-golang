package workspaces

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateWorkspaceHandler(context echo.Context) error {
	workspace := &models.Workspace{}

	println(workspace)

	if bindErr := context.Bind(workspace); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := CreateWorkspaceService(workspace); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, workspace)
	}
}

func UpdateWorkspaceHandler(context echo.Context) error {
	workspace := &models.Workspace{}

	context.Bind(workspace)

	id, _ := strconv.Atoi(context.Param("id"))

	workspace.ID = uint(id)

	if err := UpdateWorkspaceService(workspace); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, workspace)
	}
}

func ListWorkspacesHandler(context echo.Context) error {
	var workspace []models.Workspace

	if err := ListWorkspacesService(&workspace); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, workspace)
	}
}

func FindWorkspaceHandler(context echo.Context) error {
	workspace := &models.Workspace{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := FindWorkspaceService(id, workspace); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, workspace)
	}
}

func DeleteWorkspaceHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	workspace := &models.Workspace{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	if err := DeleteWorkspaceService(workspace); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
