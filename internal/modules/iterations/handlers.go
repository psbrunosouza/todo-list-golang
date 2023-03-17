package iterations

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

func CreateIterationHandler(context echo.Context) error {
	iteration := &entities.Iteration{}

	if bindErr := context.Bind(iteration); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := CreateIterationService(iteration); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, iteration)
	}
}

func UpdateIterationHandler(context echo.Context) error {
	iteration := &entities.Iteration{}

	context.Bind(iteration)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := UpdateIterationService(id, iteration); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iteration)
	}
}

func ListIterationsHandler(context echo.Context) error {
	var iterations []entities.Iteration

	if err := ListIterationsService(&iterations); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iterations)
	}
}

func FindIterationHandler(context echo.Context) error {
	iteration := &entities.Iteration{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := FindIterationService(id, iteration); err != nil {
		g_err := common.NewAppError(http.StatusNotFound, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iteration)
	}
}

func DeleteIterationHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	iteration := &entities.Iteration{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := DeleteIterationService(iteration); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
