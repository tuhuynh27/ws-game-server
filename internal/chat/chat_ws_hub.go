package chat

import (
	"encoding/json"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]string

	// Inbound messages from the clients.
	broadcast chan *Chat

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Service mongo
	service *Service
}

func NewHub(service *Service) *Hub {
	return &Hub{
		clients:    make(map[*Client]string),
		broadcast:  make(chan *Chat),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		service: 	service,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = client.user
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case newChat := <-h.broadcast:
			go h.service.InsertOneService(newChat)

			newChatJSON, _ := json.Marshal(newChat)

			for client := range h.clients {
				select {
				case client.send <- newChatJSON:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
