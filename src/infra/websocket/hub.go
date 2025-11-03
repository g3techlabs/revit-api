package websocket

import (
	"github.com/g3techlabs/revit-api/src/infra/websocket/models"
	"github.com/gofiber/contrib/websocket"
)

type Hub struct {
	clients    map[uint]*websocket.Conn
	Register   chan *models.ClientRegistration
	Unregister chan uint
	Multicast  chan *MulticastMessage
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *models.ClientRegistration),
		Unregister: make(chan uint),
		clients:    make(map[uint]*websocket.Conn),
		Multicast:  make(chan *MulticastMessage),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case registration := <-h.Register:
			h.clients[registration.ID] = registration.Conn

		case clientID := <-h.Unregister:
			if conn, ok := h.clients[clientID]; ok {
				delete(h.clients, clientID)
				conn.Close()
			}

		case msg := <-h.Multicast:
			for _, targetID := range msg.TargetUserIDs {
				if conn, ok := h.clients[targetID]; ok {
					if err := conn.WriteMessage(websocket.TextMessage, msg.Payload); err != nil {
						h.Unregister <- targetID
					}
				}
			}
		}

	}
}
