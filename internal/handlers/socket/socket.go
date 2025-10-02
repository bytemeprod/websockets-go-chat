package socket

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/bytemeprod/websockets-go-chat/internal/client"
	"github.com/bytemeprod/websockets-go-chat/internal/config"
	"github.com/bytemeprod/websockets-go-chat/internal/manager"
	"github.com/bytemeprod/websockets-go-chat/internal/tokens"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// Errors
var (
	ErrInvalidToken = errors.New("invalid token")
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Response struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewHandler(manager *manager.Manager, config config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.QueryParam("token")

		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %s\n", err.Error())
			return err
		}

		claims, err := tokens.ValidateJWT([]byte(config.SecretKey), token)
		if err != nil {
			return sendClosingResponse(conn, "error", ErrInvalidToken.Error())
		}

		sub, err := claims.GetSubject()
		if err != nil {
			return sendClosingResponse(conn, "error", ErrInvalidToken.Error())
		}

		client := client.NewClient(sub, manager, conn)
		manager.AddClient(client)

		fmt.Printf("User %s connected\n", sub)

		go client.ReadConnection()
		go client.WriteConnection()

		return nil
	}
}

func sendClosingResponse(conn *websocket.Conn, messageType, message string) error {
	response := Response{
		Type:    messageType,
		Message: message,
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal response: %s\n", err.Error())
		return err
	}
	conn.WriteMessage(websocket.BinaryMessage, data)
	return conn.Close()
}
