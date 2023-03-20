package iterations

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

type IterationHandler interface {
	Create(context echo.Context)
	List(context echo.Context)
	Get(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}

type handler struct {
	service IterationService
}

func NewIterationHandler(service IterationService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(context echo.Context) error {
	iteration := &entities.Iteration{}

	if bindErr := context.Bind(iteration); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := h.service.Create(iteration); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, iteration)
	}
}

func (h *handler) Update(context echo.Context) error {
	iteration := &entities.Iteration{}

	context.Bind(iteration)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Update(id, iteration); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iteration)
	}
}

func (h *handler) List(context echo.Context) error {
	var iterations []entities.Iteration

	if err := h.service.List(&iterations); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iterations)
	}
}

func (h *handler) Find(context echo.Context) error {
	iteration := &entities.Iteration{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Find(id, iteration); err != nil {
		g_err := common.NewAppError(http.StatusNotFound, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iteration)
	}
}

func (h *handler) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	iteration := &entities.Iteration{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := h.service.Delete(iteration); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
