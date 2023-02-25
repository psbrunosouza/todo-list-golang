package task_handlers

import (
	"net/http"
	"strconv"
	"todo-list/internal/common/error_handlers"
	task_models "todo-list/internal/modules/tasks/models"
	task_services "todo-list/internal/modules/tasks/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTaskHandler(context echo.Context) error {
	task := &task_models.Task{}

	if bindErr := context.Bind(task); bindErr != nil {
		err := error_handlers.NewGenericError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := task_services.CreateTaskService(task); serviceErr != nil {
		err := error_handlers.NewGenericError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, task)
	}
}

func UpdateTaskHandler(context echo.Context) error {
	task := &task_models.Task{}

	context.Bind(task)

	id, _ := strconv.Atoi(context.Param("id"))

	task.ID = uint(id)

	if err := task_services.UpdateTaskService(task); err != nil {
		g_err := error_handlers.NewGenericError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}

func ListTaskHandler(context echo.Context) error {
	var tasks []task_models.Task

	if err := task_services.ListTasksService(&tasks); err != nil {
		g_err := error_handlers.NewGenericError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, tasks)
	}
}

func FindTaskHandler(context echo.Context) error {
	task := &task_models.Task{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := task_services.FindTaskService(id, task); err != nil {
		g_err := error_handlers.NewGenericError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, task)
	}
}

func DeleteTaskHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	task := &task_models.Task{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	if err := task_services.DeleteTaskService(task); err != nil {
		g_err := error_handlers.NewGenericError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
