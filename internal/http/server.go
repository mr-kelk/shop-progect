package http

import (
	"errors"
	"example/shop-progect/config"
	"example/shop-progect/internal/database"
	"example/shop-progect/internal/http/handlers"
	m "example/shop-progect/internal/http/middleware"
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

	productHandler := handlers.NewProductHandler(
		service.NewProductService(
			repository.NewProductRepository(database.DB)),
	)

	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/registration", authHandler.Registration)
	authGroup.POST("/logout", authHandler.Logout)

	userGroup := e.Group("/user", m.AuthRequired)
	userGroup.GET("", handlers.GetAuthUser)

	productGroup := e.Group("/product", m.AuthRequired)
	productGroup.GET("/list", productHandler.GetProducts)
	productGroup.GET("/:uuid", productHandler.GetProductByUUID)

	productGroup.DELETE("/:uuid", productHandler.DelProductByUUID)
	productGroup.DELETE("/multiple", productHandler.DelMultipleProducts)

	productGroup.PUT("/:uuid", productHandler.UpdateProductByUUID)
	productGroup.POST("/add", productHandler.AddProduct)

	addr := net.JoinHostPort(config.Cfg.AppHost, strconv.Itoa(config.Cfg.AppPort))
	if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
