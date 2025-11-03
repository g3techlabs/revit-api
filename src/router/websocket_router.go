package router

import (
	"github.com/g3techlabs/revit-api/src/core/auth/middleware"
	"github.com/g3techlabs/revit-api/src/core/geolocation"
	ws "github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/infra/websocket/handler"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WebSocketRoute(router fiber.Router, hub *ws.Hub, geoService geolocation.IGeoLocationService, m *middleware.AuthMiddleware) {
	utils.Log.Info("WEBSOCKET route setting up...")

	webSocketHandler := handler.NewWebSocketHandler(hub, geoService)

	router.Get("/ws", m.Auth(), websocket.New(webSocketHandler.Handle))

	utils.Log.Info("WEBSOCKET route successfully set up.")
}
