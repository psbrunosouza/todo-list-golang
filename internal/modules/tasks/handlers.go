package tasks

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

func CreateTaskHandler(context echo.Context) error {
	task := &entities.Task{}

	if bindErr := context.Bind(task); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := CreateTaskService(task); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, task)
	}
}

func UpdateTaskHandler(context echo.Context) error {
	task := &entities.Task{}

	context.Bind(task)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := UpdateTaskService(id, task); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}

func ListTaskHandler(context echo.Context) error {
	var tasks []entities.Task

	if err := ListTasksService(&tasks); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {

		return context.JSON(http.StatusOK, tasks)
	}
}

func FindTaskHandler(context echo.Context) error {
	task := &entities.Task{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := FindTaskService(id, task); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}

func DeleteTaskHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	task := &entities.Task{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := DeleteTaskService(task); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}
