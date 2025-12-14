package handlers

import (
	"example/shop-progect/internal/http/validator/dto"
	"example/shop-progect/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	auth    *service.AuthService
	session *service.SessionService
}

func NewAuthHandler(auth *service.AuthService, session *service.SessionService) *AuthHandler {
	return &AuthHandler{auth: auth, session: session}
}

func (h *AuthHandler) Login(c echo.Context) error {
	req := new(dto.LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"validation_error": err.Error(),
		})
	}

	user, err := h.auth.Login(req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"login_error": err.Error()})
	}

	errSession := h.session.SetAuthUserSession(c, user)

	if errSession != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create session"})
	}

	return c.String(http.StatusOK, "Login successful")
}

func (h *AuthHandler) Registration(c echo.Context) error {
	req := new(dto.RegistrationRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"validation_error": err.Error(),
		})
	}

	user, err := h.auth.Register(req.Login, req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"registration_error": err.Error(),
		})
	}

	errSession := h.session.SetAuthUserSession(c, user)

	if errSession != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create session"})
	}

	return c.String(http.StatusCreated, "Login successful")
}

func (h *AuthHandler) Logout(c echo.Context) error {

	errSession := h.session.ClearSession(c)

	if errSession != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to clear session"})
	}

	return c.String(http.StatusOK, "Logout successful")
}
