package socket

import (
	"log"

	"github.com/bytemeprod/websockets-go-chat/internal/client"
	"github.com/bytemeprod/websockets-go-chat/internal/types"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewHandler(manager types.Manager) echo.HandlerFunc {
	return func(с echo.Context) error {
		conn, err := upgrader.Upgrade(с.Response(), с.Request(), nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return err
		}
		defer conn.Close()

		client := client.NewClient(manager, conn)
		manager.AddClient(client)

		go client.ReadConnection()
		go client.WriteConnection()

		return nil
	}
}
