package model

type Hub struct {
	// Registered connections.
	connections map[*Connection]bool

	// Inbound messages from the Connections.
	broadcast chan interface{}

	// Register requests from the Connections.
	register chan *Connection

	// Unregister requests from Connections.
	unregister chan *Connection
}

func NewHub() *Hub {
	return &Hub{
		broadcast:   make(chan interface{}),
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
		connections: make(map[*Connection]bool),
	}
}

func (h *Hub) RegisterConnection(conn *Connection) {
	h.register <- conn
}

func (h *Hub) UnregisterConnection(conn *Connection) {
	h.unregister <- conn
}

func (h *Hub) BroadcastMessage(message interface{}) {
	h.broadcast <- message
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}
