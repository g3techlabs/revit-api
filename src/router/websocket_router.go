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

func WebSocketRoute(router fiber.Router, hub *ws.Hub, geoService geolocation.IGeoLocationService, m *middleware.AuthMiddleware, logger utils.ILogger) {
	logger.Info("WEBSOCKET route setting up...")

	webSocketHandler := handler.NewWebSocketHandler(hub, geoService, logger)

	router.Get("/ws", m.Auth(), websocket.New(webSocketHandler.Handle))

	logger.Info("WEBSOCKET route successfully set up.")
}
