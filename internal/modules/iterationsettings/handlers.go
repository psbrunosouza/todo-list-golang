package iterationsettings

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/models"

	"github.com/labstack/echo/v4"
)

func CreateIterationSettingHandler(context echo.Context) error {
	iterationSetting := &models.IterationSetting{}

	if bindErr := context.Bind(iterationSetting); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := CreateIterationSettingService(iterationSetting); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, iterationSetting)
	}
}

func UpdateIterationSettingHandler(context echo.Context) error {
	iterationSetting := &models.IterationSetting{}

	context.Bind(iterationSetting)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := UpdateIterationSettingService(id, iterationSetting); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iterationSetting)
	}
}

func ListIterationSettingsHandler(context echo.Context) error {
	var iterationSettings []models.IterationSetting

	if err := ListIterationSettingsService(&iterationSettings); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iterationSettings)
	}
}

func FindIterationSettingHandler(context echo.Context) error {
	iterationSetting := &models.IterationSetting{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := FindIterationSettingService(id, iterationSetting); err != nil {
		g_err := common.NewAppError(http.StatusNotFound, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iterationSetting)
	}
}

func DeleteIterationSettingHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	iterationSetting := &models.IterationSetting{
		Default: models.Default{
			ID: uint(id),
		},
	}

	if err := DeleteIterationSettingService(iterationSetting); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
