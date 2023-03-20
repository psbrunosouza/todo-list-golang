package ratings

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

type RatingHandler interface {
	Create(context echo.Context)
	List(context echo.Context)
	Get(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}

type handler struct {
	service RatingService
}

func NewRatingHandler(service RatingService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(context echo.Context) error {
	rating := &entities.Rating{}

	if bindErr := context.Bind(rating); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := h.service.Create(rating); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, rating)
	}
}

func (h *handler) Update(context echo.Context) error {
	rating := &entities.Rating{}

	context.Bind(rating)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Update(id, rating); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, rating)
	}
}

func (h *handler) List(context echo.Context) error {
	var ratings []entities.Rating

	if err := h.service.List(&ratings); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, ratings)
	}
}

func (h *handler) Find(context echo.Context) error {
	rating := &entities.Rating{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Find(id, rating); err != nil {
		g_err := common.NewAppError(http.StatusNotFound, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, rating)
	}
}

func (h *handler) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	rating := &entities.Rating{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := h.service.Delete(rating); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, &struct{}{})
	}
}
