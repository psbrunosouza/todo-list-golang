package task_handlers

import (
	"net/http"
	"strconv"
	errorhandlers "todo-list/internal/common/error_handlers"
	tasksService "todo-list/internal/modules/tasks/services"
	taskModels "todo-list/internal/modules/tasks/task_models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTaskHandler(context echo.Context) error {
	task := &taskModels.Task{}

	if bindErr := context.Bind(task); bindErr != nil {
		err := errorhandlers.NewGenericError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := tasksService.CreateTaskService(task); serviceErr != nil {
		err := errorhandlers.NewGenericError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, task)
	}
}

func UpdateTaskHandler(context echo.Context) error {
	task := &taskModels.Task{}

	context.Bind(task)

	id, _ := strconv.Atoi(context.Param("id"))

	task.ID = uint(id)

	if err := tasksService.UpdateTaskService(task); err != nil {
		g_err := errorhandlers.NewGenericError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}

func ListTaskHandler(context echo.Context) error {
	var tasks []taskModels.Task

	if err := tasksService.ListTasksService(&tasks); err != nil {
		g_err := errorhandlers.NewGenericError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, tasks)
	}
}

func FindTaskHandler(context echo.Context) error {
	task := &taskModels.Task{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := tasksService.FindTaskService(id, task); err != nil {
		g_err := errorhandlers.NewGenericError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}

func DeleteTaskHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	task := &taskModels.Task{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	if err := tasksService.DeleteTaskService(task); err != nil {
		g_err := errorhandlers.NewGenericError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
