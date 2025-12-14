package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAuthUser(c echo.Context) error {
	return c.String(http.StatusOK, "GetAuthUser")
}
