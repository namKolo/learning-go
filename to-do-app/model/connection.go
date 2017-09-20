package model

import (
	"github.com/gorilla/websocket"
)

type Connection struct {
	ws   *websocket.Conn
	send chan interface{}
}

func NewConnection(send chan interface{}, ws *websocket.Conn) *Connection {
	return &Connection{
		ws:   ws,
		send: send,
	}
}

func (c *Connection) Read() {
	for {
		_, _, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func (c *Connection) Write() {
	for change := range c.send {
		err := c.ws.WriteJSON(change)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
