package handler

import (
	"encoding/json"
	"fmt"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	geoservice "github.com/g3techlabs/revit-api/src/core/geolocation/service"
	routeservice "github.com/g3techlabs/revit-api/src/core/route/service"
	ws "github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/infra/websocket/input"
	"github.com/g3techlabs/revit-api/src/infra/websocket/models"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/contrib/websocket"
)

type WebSocketHandler struct {
	hub          *ws.Hub
	geoService   geoservice.IGeoLocationService
	routeService routeservice.IRouteService
	logger       utils.ILogger
}

func NewWebSocketHandler(hub *ws.Hub, geoService geoservice.IGeoLocationService, routeService routeservice.IRouteService, logger utils.ILogger) *WebSocketHandler {
	return &WebSocketHandler{
		hub:          hub,
		geoService:   geoService,
		routeService: routeService,
		logger:       logger,
	}
}

func (h *WebSocketHandler) Handle(c *websocket.Conn) {
	userId, err := h.getUserId(c)
	if err != nil {
		c.Close()
		return
	}

	h.registerClientInHub(c, userId)

	defer h.unregisterClientInHub(userId)

	for {
		var message ws.WebSocketMessage
		if err := c.ReadJSON(&message); err != nil {
			h.logger.Errorf("Error reading WebSocketMessage: %v", err)
			break
		}

		switch message.Event {
		case "put-user-location":
			var payload geoinput.Coordinates
			if err := json.Unmarshal(message.Payload, &payload); err != nil {
				continue
			}

			if err := h.geoService.PutUserLocation(userId, &payload); err != nil {
				continue
			}
		case "start-route":
			var payload input.StartRoutePayload
			if err := json.Unmarshal(message.Payload, &payload); err != nil {
				continue
			}

			if err := h.routeService.StartRoute(userId, payload.RouteID); err != nil {
				h.logger.Errorf("Error starting route %d: %v", payload.RouteID, err)
				continue
			}
		case "participant-finish-route":
			var payload geoinput.Coordinates
			if err := json.Unmarshal(message.Payload, &payload); err != nil {
				continue
			}

			if err := h.routeService.FinishRouteParticipant(userId, &payload); err != nil {
				h.logger.Errorf("Error finishing route participant %d: %v", userId, err)
				continue
			}
		case "leave-route":
			var payload geoinput.Coordinates

			if err := json.Unmarshal(message.Payload, &payload); err != nil {
				continue
			}

			if err := h.routeService.LeaveRoute(userId, &payload); err != nil {
				h.logger.Errorf("Error leaving route participant %d: %v", userId, err)
				continue
			}
		}
	}
}

func (h *WebSocketHandler) getUserId(c *websocket.Conn) (uint, error) {
	userId, ok := c.Locals("userId").(uint)
	if !ok {
		return 0, fmt.Errorf("invalid userID")
	}

	return userId, nil
}

func (h *WebSocketHandler) registerClientInHub(c *websocket.Conn, userId uint) {
	registrationData := &models.ClientRegistration{
		ID:   userId,
		Conn: c,
	}

	h.hub.Register <- registrationData
}

func (h *WebSocketHandler) unregisterClientInHub(userId uint) {
	if err := h.geoService.RemoveUserLocation(userId); err != nil {
		return
	}

	h.hub.Unregister <- userId
}
