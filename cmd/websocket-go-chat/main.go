package main

import (
	"github.com/bytemeprod/websockets-go-chat/internal/config"
	"github.com/labstack/echo"
)

func main() {
	config := config.MustLoadConfig()
	e := echo.New()
	//e.Logger.SetLevel(log.OFF)
	e.Static("/", "./frontend")
	//e.GET("/ws", handlers.ServeWs)
	e.Start(config.Host + ":" + config.Port)
}
