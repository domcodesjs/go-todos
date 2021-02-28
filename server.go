package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/", GetTodos)

	e.Logger.Fatal(e.Start(":5000"))
}


