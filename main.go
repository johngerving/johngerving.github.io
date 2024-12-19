package main

import (
	"errors"
	"log/slog"
	"net/http"

	"johngerving/johngerving.github.io/pkg/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", handlers.IndexHandler)

	if err := e.Start(":8080"); err != nil && errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
