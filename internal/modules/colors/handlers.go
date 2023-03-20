package colors

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

type ColorHandler interface {
	Create(context echo.Context)
	List(context echo.Context)
	Get(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}

type handler struct {
	service ColorService
}

func NewColorHandler(service ColorService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(context echo.Context) error {
	color := &entities.Color{}

	if bindErr := context.Bind(color); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := h.service.Create(color); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, color)
	}
}

func (h *handler) Update(context echo.Context) error {
	color := &entities.Color{}

	context.Bind(color)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Update(id, color); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, color)
	}
}

func (h *handler) List(context echo.Context) error {
	var colors []entities.Color

	if err := h.service.List(&colors); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, colors)
	}
}

func (h *handler) Find(context echo.Context) error {
	color := &entities.Color{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Find(id, color); err != nil {
		g_err := common.NewAppError(http.StatusNotFound, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, color)
	}
}

func (h *handler) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	color := &entities.Color{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := h.service.Delete(color); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
