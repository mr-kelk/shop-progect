package middleware

import (
	"example/shop-progect/internal/model"
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

		id, okID := sess.Values["auth_id"].(string)
		login, okLogin := sess.Values["auth_login"].(string)
		email, okEmail := sess.Values["auth_email"].(string)
		role, okRole := sess.Values["auth_role_name"].(string)

		if !okID || !okLogin || !okEmail || !okRole {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		authUser := &model.UserSess{
			ID:       id,
			Login:    login,
			Email:    email,
			RoleName: role,
		}

		c.Set("authUser", authUser)

		return next(c)
	}
}
