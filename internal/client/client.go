package client

import (
	"fmt"

	"github.com/bytemeprod/websockets-go-chat/internal/manager"
	"github.com/gorilla/websocket"
)

type Client struct {
	manager *manager.Manager
	conn    *websocket.Conn
	egress  chan []byte
}

func NewClient(manager *manager.Manager, conn *websocket.Conn) *Client {
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
		_, msg, err := c.conn.ReadMessage() // just read message now
		if err != nil {
			fmt.Printf("Error reading message: %v", err)
			break
		}
		fmt.Printf("Message: %s\n", string(msg))
		for cl := range c.manager.Clients {
			cl.AddToEgress(msg)
		}
	}
}

func (c *Client) AddToEgress(message []byte) {
	c.egress <- message
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
		default:
			// ...
		}
	}
}
