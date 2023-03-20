package tasks

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	Create(context echo.Context)
	List(context echo.Context)
	Get(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}

type handler struct {
	service TaskService
}

func NewTaskHandler(service TaskService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(context echo.Context) error {
	task := &entities.Task{}

	if bindErr := context.Bind(task); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := h.service.Create(task); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, task)
	}
}

func (h *handler) Update(context echo.Context) error {
	task := &entities.Task{}

	context.Bind(task)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Update(id, task); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}

func (h *handler) List(context echo.Context) error {
	var tasks []entities.Task

	if err := h.service.List(&tasks); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {

		return context.JSON(http.StatusOK, tasks)
	}
}

func (h *handler) Find(context echo.Context) error {
	task := &entities.Task{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Find(id, task); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}

func (h *handler) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	task := &entities.Task{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := h.service.Delete(task); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}
