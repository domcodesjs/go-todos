package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	CreatedAt time.Time `json:"CreatedAt"`
	Text string `json:"Text"`
}


func GetTodos(c echo.Context) error {
	todos := []Todo{
		{
			CreatedAt:  time.Now(),
			Text: "todo 1",
		},
		{
			CreatedAt:  time.Now(),
			Text: "todo 2",
		},
		{
			CreatedAt:  time.Now(),
			Text: "todo 3",
		},
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return c.JSON(http.StatusOK, todos)
}
