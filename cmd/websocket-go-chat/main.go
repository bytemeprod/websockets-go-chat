package main

import (
	"context"
	"log"

	"github.com/bytemeprod/websockets-go-chat/internal/config"
	"github.com/bytemeprod/websockets-go-chat/internal/handlers/login"
	"github.com/bytemeprod/websockets-go-chat/internal/handlers/socket"
	"github.com/bytemeprod/websockets-go-chat/internal/manager"
	"github.com/bytemeprod/websockets-go-chat/internal/redisstore"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config := config.MustLoadConfig()
	storage, err := redisstore.NewClient(config.RedisConfig.Addr, config.RedisConfig.Password)
	if err != nil {
		log.Fatal(err)
	}
	manager := manager.NewManager(storage)
	ctx := context.Background()

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/ws", socket.NewHandler(manager))

	e.Static("/", "./frontend")
	e.POST("/login", login.NewHandler(ctx, storage))

	e.Start(config.Host + ":" + config.Port)
}
