package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		authUser := sess.Values["auth"]
		if authUser == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		c.Set("authUser", authUser)

		return next(c)
	}
}
