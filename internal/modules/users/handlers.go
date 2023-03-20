package users

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Create(context echo.Context)
	List(context echo.Context)
	Get(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}

type handler struct {
	service UserService
}

func NewUserHandler(service UserService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(context echo.Context) error {
	user := &entities.User{}

	if bindErr := context.Bind(user); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := h.service.Create(user); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, user)
	}
}

func (h *handler) Update(context echo.Context) error {
	user := &entities.User{}

	context.Bind(user)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Update(id, user); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, user)
	}
}

func (h *handler) List(context echo.Context) error {
	var users []entities.User

	if err := h.service.List(&users); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, users)
	}
}

func (h *handler) Find(context echo.Context) error {
	user := &entities.User{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := h.service.Find(id, user); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, user)
	}
}

func (h *handler) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	user := &entities.User{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := h.service.Delete(user); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, &struct{}{})
	}
}
