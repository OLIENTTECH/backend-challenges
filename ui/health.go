package ui

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health interface {
	GetHealth(c echo.Context) error
}

type health struct{}

func NewHealth() Health {
	return &health{}
}

func (h *health) GetHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
