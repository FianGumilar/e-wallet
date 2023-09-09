package main

import (
	"fmt"

	"github.com/FianGumilar/e-wallet/config"
	"github.com/FianGumilar/e-wallet/infrastructure/database"
	"github.com/FianGumilar/e-wallet/infrastructure/http/handler"
	"github.com/FianGumilar/e-wallet/repository"
	"github.com/FianGumilar/e-wallet/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.NewAppConfig()

	cacheConnection, err := repository.NewRedisRepository(conf)
	if err != nil {
		fmt.Printf("Error create cache connection: %s", err)
	}
	dbPgConnection := database.GetDbPgConnection(conf)

	// repository
	userRepository := repository.NewUserRepository(dbPgConnection)

	// service
	userService := service.NewUserService(userRepository, cacheConnection)

	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// handler
	handler.NewAuthHandler(e, userService)

	e.Start(":8080")
}
