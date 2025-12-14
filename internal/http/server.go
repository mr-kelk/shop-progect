package http

import (
	"errors"
	"example/shop-progect/config"
	"example/shop-progect/internal/database"
	"example/shop-progect/internal/http/handlers"
	"example/shop-progect/internal/http/validator"
	"example/shop-progect/internal/repository"
	"example/shop-progect/internal/service"
	"net"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartHttp() {
	e := echo.New()
	redisStore := database.NewRedisStore([]byte(config.Cfg.SessionKey))
	defer redisStore.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(session.Middleware(redisStore))
	e.Validator = validator.NewValidator()

	e.GET("/healthcheck", handlers.Healthcheck)

	authHandler := handlers.NewAuthHandler(
		service.NewAuthService(
			repository.NewUserRepository(database.DB),
		),
		service.NewSessionService(),
	)

	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/registration", authHandler.Registration)
	authGroup.POST("/logout", authHandler.Logout)

	userGroup := e.Group("/user")
	userGroup.GET("", handlers.GetAuthUser)

	productGroup := e.Group("/product")
	productGroup.GET("/list", handlers.GetProducts)
	productGroup.GET("/:uuid", handlers.GetProductByUUID)

	addr := net.JoinHostPort(config.Cfg.AppHost, strconv.Itoa(config.Cfg.AppPort))

	if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
