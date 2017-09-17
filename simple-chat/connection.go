package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan interface{}
}

// wait for changes
func (c *connection) writer() {
	for change := range c.send {
		err := c.ws.WriteJSON(change)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func wsHandler(h hub) http.HandlerFunc {
	log.Println("Starting websocket server")
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade to websocket
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		// create new connection and send to register channel of hub
		c := &connection{send: make(chan interface{}, 256), ws: ws}
		h.register <- c
		numOfMessage := NumOfMessage{Count: len(h.connections) + 1}
		h.broadcast <- numOfMessage

		defer func() {
			h.unregister <- c
			numOfMessage := NumOfMessage{Count: len(h.connections) - 1}
			h.broadcast <- numOfMessage
		}()
		go c.writer()
		// read message from client  - then broadcast it
		for {
			var msg Message
			err := c.ws.ReadJSON(&msg)

			if err != nil {
				break
			}
			h.broadcast <- msg
		}
	}
}
