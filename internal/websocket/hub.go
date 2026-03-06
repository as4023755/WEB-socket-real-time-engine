package websocket

type Hub struct {
	clients    map[*Client]bool
	rooms      map[string]map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		rooms:      make(map[string]map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}
func (h *Hub) Run() {

	for {
		select {

		case client := <-h.register:

			h.clients[client] = true

			room := client.room

			if _, ok := h.rooms[room]; !ok {
				h.rooms[room] = make(map[*Client]bool)
			}

			h.rooms[room][client] = true

		case client := <-h.unregister:

			delete(h.clients, client)

			room := client.room

			if clients, ok := h.rooms[room]; ok {

				delete(clients, client)

				if len(clients) == 0 {
					delete(h.rooms, room)
				}
			}

		case message := <-h.broadcast:

			if clients, ok := h.rooms[message.Room]; ok {

				for client := range clients {
					client.send <- message.Data
				}

			}

		}
	}
}
