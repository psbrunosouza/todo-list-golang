package subtasks

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/models"

	"github.com/labstack/echo/v4"
)

func CreateSubTaskHandler(context echo.Context) error {
	subtask := &models.SubTask{}

	if bindErr := context.Bind(subtask); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := CreateSubTaskService(subtask); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, subtask)
	}
}

func UpdateSubTaskHandler(context echo.Context) error {
	subtask := &models.SubTask{}

	context.Bind(subtask)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := UpdateSubTaskService(id, subtask); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, subtask)
	}
}

func DeleteSubTaskHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	subtask := &models.SubTask{
		Default: models.Default{
			ID: uint(id),
		},
	}

	if err := DeleteSubTaskService(subtask); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}

func FindSubTaskHandler(context echo.Context) error {
	subtask := &models.SubTask{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := FindSubTaskService(id, subtask); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, subtask)
	}
}

func ListSubTasksHandler(context echo.Context) error {
	var subtasks []models.SubTask

	if err := ListSubTasksService(&subtasks); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, subtasks)
	}
}

func MarkSubtaskAsDoneHandler(context echo.Context) error {
	subtask := &models.SubTask{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := MarkAsDoneService(id, subtask); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, subtask)
	}
}
