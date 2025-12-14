package app

import (
	"example/shop-progect/config"
	"example/shop-progect/internal/database"
	"example/shop-progect/internal/http"
)

func init() {
	config.NewConfig()
}

func Run() {
	if err := database.Init(); err != nil {
		panic(err)
	}

	defer database.DB.Close()

	http.StartHttp()
}
