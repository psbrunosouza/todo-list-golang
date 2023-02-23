package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
}

func main() {
	var tasks []Task = []Task{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		response, _ := json.Marshal(tasks)

		return c.JSON(http.StatusOK, string(response))
	})
	e.POST("/", func(context echo.Context) error {
		var task Task

		bindError := context.Bind(&task)

		if bindError != nil {
			context.String(http.StatusBadRequest, "bad request")
		}

		tasks = append(tasks, task)

		response, _ := json.Marshal(tasks)
		return context.JSON(http.StatusOK, string(response))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
