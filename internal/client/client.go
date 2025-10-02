package client

import (
	"fmt"
	"log"
	"time"

	"github.com/bytemeprod/websockets-go-chat/internal/manager"
	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
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

	c.conn.SetPongHandler(c.pongHalder)

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		fmt.Printf("Failed to set read deadline: %v\n", err)
		return
	}

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

func (c *Client) pongHalder(pongMsg string) error {
	log.Println("Pong message received")
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}

func (c *Client) AddToEgress(message []byte) {
	c.egress <- message
}

func (c *Client) WriteConnection() {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()
	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, message); err != nil {
					fmt.Printf("Connection closed: %v", err)
				}
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Printf("Error writing message: %v", err)
			}
		case <-ticker.C:
			log.Println("Ping message send")
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Printf("Error writing ping message: %v", err)
			}
		}
	}
}

func (c *Client) GetUsername() string {
	return c.username
}
