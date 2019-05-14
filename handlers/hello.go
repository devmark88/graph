package handlers

import (
	"net/http"

	"github.com/devmark88/unireg/config"

	"github.com/labstack/echo"
)

// Hello => hello world
func Hello(c echo.Context, app *config.AppContext) error {
	return c.String(http.StatusOK, "Hello, World!")
}
