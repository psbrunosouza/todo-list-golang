package iterationsettings

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

type IterationSettingHandler interface {
	Create(context echo.Context)
	List(context echo.Context)
	Get(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}

type handler struct {
	service IterationSettingService
}

func NewIterationSettingHandler(service IterationSettingService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(context echo.Context) error {
	iterationSetting := &entities.IterationSetting{}

	if bindErr := context.Bind(iterationSetting); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := h.service.Create(iterationSetting); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, iterationSetting)
	}
}

func (h *handler) Update(context echo.Context) error {
	iterationSetting := &entities.IterationSetting{}

	context.Bind(iterationSetting)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Update(id, iterationSetting); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iterationSetting)
	}
}

func (h *handler) List(context echo.Context) error {
	var iterationSettings []entities.IterationSetting

	if err := h.service.List(&iterationSettings); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iterationSettings)
	}
}

func (h *handler) Find(context echo.Context) error {
	iterationSetting := &entities.IterationSetting{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Find(id, iterationSetting); err != nil {
		g_err := common.NewAppError(http.StatusNotFound, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, iterationSetting)
	}
}

func (h *handler) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	iterationSetting := &entities.IterationSetting{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := h.service.Delete(iterationSetting); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
