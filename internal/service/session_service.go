package service

import (
	"example/shop-progect/internal/model"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type SessionService struct{}

func NewSessionService() *SessionService {
	return &SessionService{}
}

func (s *SessionService) SetAuthUserSession(c echo.Context, user *model.UserSess) error {
	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values["auth_id"] = user.ID
	sess.Values["auth_login"] = user.Login
	sess.Values["auth_email"] = user.Email
	sess.Values["auth_role_name"] = user.RoleName

	return sess.Save(c.Request(), c.Response())
}

func (s *SessionService) ClearSession(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	return sess.Save(c.Request(), c.Response())
}
