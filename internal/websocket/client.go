package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn

	send chan []byte

	username string
	room     string
}

func (c *Client) readPump() {

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {

		_, message, err := c.conn.ReadMessage()

		if err != nil {
			break
		}

		c.hub.broadcast <- Message{
			Room: c.room,
			Data: message,
		}

	}
}
func (c *Client) writePump() {

	for {

		message := <-c.send

		c.conn.WriteMessage(websocket.TextMessage, message)

	}
}
