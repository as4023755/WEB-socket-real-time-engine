package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) readPump() {
	fmt.Println("READ PUMP STARTED")
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("Reading message...")
		fmt.Println("Received:", string(message))
		c.hub.broadcast <- message
	}
}
func (c *Client) writePump() {
	for {
		message := <-c.send
		c.conn.WriteMessage(websocket.TextMessage, message)
	}
}
