package workspaces

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

type WorkspaceHandler interface {
	Create(context echo.Context)
	List(context echo.Context)
	Get(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}

type handler struct {
	service WorkspaceService
}

func NewWorkspaceHandler(service WorkspaceService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(context echo.Context) error {
	workspace := &entities.Workspace{}

	println(workspace)

	if bindErr := context.Bind(workspace); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := h.service.Create(workspace); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, workspace)
	}
}

func (h *handler) Update(context echo.Context) error {
	workspace := &entities.Workspace{}

	context.Bind(workspace)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Update(id, workspace); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, workspace)
	}
}

func (h *handler) List(context echo.Context) error {
	var workspace []entities.Workspace

	if err := h.service.List(&workspace); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, workspace)
	}
}

func (h *handler) Find(context echo.Context) error {
	workspace := &entities.Workspace{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Find(id, workspace); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, workspace)
	}
}

func (h *handler) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	workspace := &entities.Workspace{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := h.service.Delete(workspace); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
