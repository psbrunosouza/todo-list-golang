package users

import (
	"net/http"
	"strconv"
	"todo-list/internal/common"
	"todo-list/internal/entities"

	"github.com/labstack/echo/v4"
)

func CreateUsersHandler(context echo.Context) error {
	user := &entities.User{}

	if bindErr := context.Bind(user); bindErr != nil {
		err := common.NewAppError(http.StatusBadRequest, bindErr)
		return context.JSON(http.StatusBadRequest, err)
	}

	if serviceErr := CreateUserService(user); serviceErr != nil {
		err := common.NewAppError(http.StatusBadRequest, serviceErr)
		return context.JSON(http.StatusBadRequest, err)
	} else {
		return context.JSON(http.StatusCreated, user)
	}
}

func UpdateUserHandler(context echo.Context) error {
	user := &entities.User{}

	context.Bind(user)

	id, _ := strconv.Atoi(context.Param("id"))

	if err := UpdateUserService(id, user); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, user)
	}
}

func ListUsersHandler(context echo.Context) error {
	var users []entities.User

	if err := ListUsersService(&users); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, users)
	}
}

func FindUserHandler(context echo.Context) error {
	user := &entities.User{}

	id, _ := strconv.Atoi(context.Param("id"))

	if err := FindUserService(id, user); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, "Record not found")
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, user)
	}
}

func DeleteUserHandler(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))

	user := &entities.User{
		Default: entities.Default{
			ID: uint(id),
		},
	}

	if err := DeleteUserService(user); err != nil {
		g_err := common.NewAppError(http.StatusBadRequest, err)
		return context.JSON(http.StatusBadRequest, g_err)
	} else {
		return context.JSON(http.StatusOK, string("{}"))
	}
}
