package handlers

import (
	"example/shop-progect/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAuthUser(c echo.Context) error {
	userAny := c.Get("authUser").(*model.UserSess)
	return c.JSON(http.StatusOK, userAny)
}
