package tasks

import (
	"net/http"
	taskModels "todo-list/internal/modules/tasks/models"
	taskServices "todo-list/internal/modules/tasks/services"

	"github.com/labstack/echo/v4"
)

func CreateTaskHandler(context echo.Context) error {
	task := &taskModels.Task{}

	if error := context.Bind(task); error != nil {
		panic("error")
	}

	taskServices.CreateTaskService(task)

	return context.JSON(http.StatusCreated, task)
}

// func UpdateTaskHandler(context echo.Context) error {

// }

// func ListTaskHandler(context echo.Context) error {

// }

// func FindTaskHandler(context echo.Context) error {

// }

// func DeleteTaskHandler(context echo.Context) error {

// }
