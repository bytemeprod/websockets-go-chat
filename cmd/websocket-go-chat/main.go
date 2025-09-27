package main

import (
	"github.com/bytemeprod/websockets-go-chat/internal/config"
	"github.com/bytemeprod/websockets-go-chat/internal/handlers/socket"
	"github.com/bytemeprod/websockets-go-chat/internal/manager"
	"github.com/labstack/echo"
)

func main() {
	config := config.MustLoadConfig()
	manager := manager.NewManager()
	e := echo.New()
	e.Static("/", "./frontend")
	e.GET("/ws", socket.NewHandler(manager))
	e.Start(config.Host + ":" + config.Port)
}
