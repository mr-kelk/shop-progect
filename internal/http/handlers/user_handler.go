package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAuthUser(c echo.Context) error {
	user := c.Get("authUser")
	return c.JSON(http.StatusOK, user)
}
