package client

import (
	"fmt"

	"github.com/bytemeprod/websockets-go-chat/internal/manager"
	"github.com/gorilla/websocket"
)

type Client struct {
	username string
	manager  *manager.Manager
	conn     *websocket.Conn
	egress   chan []byte
}

func NewClient(username string, manager *manager.Manager, conn *websocket.Conn) *Client {
	return &Client{
		username: username,
		manager:  manager,
		conn:     conn,
		egress:   make(chan []byte),
	}
}

func (c *Client) ReadConnection() {
	defer func() {
		fmt.Printf("user %s disconnected\n", c.username)
		c.conn.Close()
		close(c.egress)
		c.manager.RemoveClient(c)
	}()
	for {
		_, msg, err := c.conn.ReadMessage() // just read message now
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		msgWithuser := c.username + ": " + string(msg)

		for cl := range c.manager.Clients {
			cl.AddToEgress([]byte(msgWithuser))
		}
	}
}

func (c *Client) AddToEgress(message []byte) {
	c.egress <- message
}

func (c *Client) WriteConnection() {
	for message := range c.egress {
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Printf("Error writing message: %v", err)
		}
	}
}

func (c *Client) GetUsername() string {
	return c.username
}
