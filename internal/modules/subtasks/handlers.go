package subtasks

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

type SubTaskHandler interface {
	Create(context echo.Context)
	List(context echo.Context)
	Get(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}

type handler struct {
	service SubTaskService
}

func NewSubTaskHandler(service SubTaskService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(context echo.Context) error {
	subtask := &entities.SubTask{}

	if bindErr := context.Bind(subtask); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := h.service.Create(subtask); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, subtask)
	}
}

func (h *handler) Update(context echo.Context) error {
	subtask := &entities.SubTask{}

	context.Bind(subtask)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Update(id, subtask); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, subtask)
	}
}

func (h *handler) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	subtask := &entities.SubTask{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := h.service.Delete(subtask); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}

func (h *handler) Find(context echo.Context) error {
	subtask := &entities.SubTask{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Find(id, subtask); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, subtask)
	}
}

func (h *handler) List(context echo.Context) error {
	var subtasks []entities.SubTask

	if err := h.service.List(&subtasks); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, subtasks)
	}
}
