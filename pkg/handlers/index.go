package handlers

import (
	"context"
	"johngerving/johngerving.github.io/pkg/views"

	"github.com/labstack/echo/v4"
)

// GET /
func IndexHandler(c echo.Context) error {
	return views.Index().Render(context.Background(), c.Response().Writer)
}
