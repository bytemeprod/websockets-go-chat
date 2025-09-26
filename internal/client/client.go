package client

import (
	"fmt"

	"github.com/bytemeprod/websockets-go-chat/internal/types"
	"github.com/gorilla/websocket"
)

type Client struct {
	manager types.Manager
	conn    *websocket.Conn
	egress  chan []byte
}

func NewClient(manager types.Manager, conn *websocket.Conn) *Client {
	return &Client{
		manager: manager,
		conn:    conn,
		egress:  make(chan []byte),
	}
}

func (c *Client) ReadConnection() {
	defer func() {
		c.conn.Close()
		c.manager.RemoveClient(c)
	}()
	for {
		_, _, err := c.conn.ReadMessage() // just read message now
		if err != nil {
			fmt.Printf("Error reading message: %v", err)
			break
		}
	}
}

func (c *Client) WriteConnection() {
	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				fmt.Printf("Failed to write message, egress chan is closed")
				if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					fmt.Printf("Failed to close connection: %v", err)
				}
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Printf("Error writing message: %v", err)
			}
		}
	}
}
