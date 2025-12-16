package middleware

import (
	"example/shop-progect/internal/model"
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

func RoleRequired(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authUser, ok := c.Get("authUser").(*model.UserSess)
			if !ok || authUser == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			if allowed := slices.Contains(roles, authUser.RoleName); !allowed {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden")
			}

			return next(c)
		}
	}
}
