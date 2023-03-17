package ratings

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

func CreateRatingHandler(context echo.Context) error {
	rating := &entities.Rating{}

	if bindErr := context.Bind(rating); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := CreateRatingService(rating); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, rating)
	}
}

func UpdateRatingHandler(context echo.Context) error {
	rating := &entities.Rating{}

	context.Bind(rating)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := UpdateRatingService(id, rating); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, rating)
	}
}

func ListRatingsHandler(context echo.Context) error {
	var ratings []entities.Rating

	if err := ListRatingsService(&ratings); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, ratings)
	}
}

func FindRatingHandler(context echo.Context) error {
	rating := &entities.Rating{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := FindRatingService(id, rating); err != nil {
		g_err := common.NewAppError(http.StatusNotFound, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, rating)
	}
}

func DeleteRatingHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	rating := &entities.Rating{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := DeleteRatingService(rating); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, &struct{}{})
	}
}
