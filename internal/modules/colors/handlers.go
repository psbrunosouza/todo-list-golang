package colors

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

func CreateColorHandler(context echo.Context) error {
	color := &entities.Color{}

	if bindErr := context.Bind(color); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := CreateColorService(color); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, color)
	}
}

func UpdateColorHandler(context echo.Context) error {
	color := &entities.Color{}

	context.Bind(color)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := UpdateColorService(id, color); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, color)
	}
}

func ListColorsHandler(context echo.Context) error {
	var colors []entities.Color

	if err := ListColorsService(&colors); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, colors)
	}
}

func FindColorHandler(context echo.Context) error {
	color := &entities.Color{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := FindColorService(id, color); err != nil {
		g_err := common.NewAppError(http.StatusNotFound, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, color)
	}
}

func DeleteColorHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	color := &entities.Color{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := DeleteColorService(color); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
