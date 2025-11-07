package websocket

import (
	"encoding/json"

	"github.com/g3techlabs/revit-api/src/infra/websocket/models"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/contrib/websocket"
)

type Hub struct {
	clients    map[uint]*websocket.Conn
	Register   chan *models.ClientRegistration
	Unregister chan uint
	Multicast  chan *MulticastMessage
	logger     utils.ILogger
}

func NewHub(logger utils.ILogger) *Hub {
	return &Hub{
		Register:   make(chan *models.ClientRegistration),
		Unregister: make(chan uint),
		clients:    make(map[uint]*websocket.Conn),
		Multicast:  make(chan *MulticastMessage),
		logger:     logger,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case registration := <-h.Register:
			h.clients[registration.ID] = registration.Conn
			h.logger.Infof("New client registered in Hub: %d", registration.ID)

		case clientID := <-h.Unregister:
			if conn, ok := h.clients[clientID]; ok {
				delete(h.clients, clientID)
				conn.Close()
				h.logger.Infof("Unregistered client from Hub: %d", clientID)
			} else {
				h.logger.Errorf("Error in unregister client operation: Client %d not found", clientID)
			}

		case msg := <-h.Multicast:
			for _, targetID := range msg.TargetUserIDs {
				if conn, ok := h.clients[targetID]; ok {
					if err := conn.WriteMessage(websocket.TextMessage, msg.Payload); err != nil {
						h.logger.Errorf("Error in WriteMessage Operation to Client %d: %v. Unregistering...", targetID, err)
						h.Unregister <- targetID
					}
				} else {
					h.logger.Warnf("Multicast: Target client %d not found. Message not sent.", targetID)
				}
			}
		}
	}
}

func (h *Hub) SendMulticastMessage(targetIds []uint, payload any) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	multicastMessage := MulticastMessage{
		TargetUserIDs: targetIds,
		Payload:       payloadBytes,
	}

	h.Multicast <- &multicastMessage

	return nil
}
