package router

import (
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/event/controller"
	"github.com/g3techlabs/revit-api/core/event/service"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
)

func EventRoutes(router fiber.Router, eventService service.IEventService, m *middleware.AuthMiddleware) {
	utils.Log.Info("Setting up EVENT routes...")

	eventController := controller.NewEventController(eventService)

	event := router.Group("/event", m.Auth())
	event.Post("/", eventController.CreateEvent)
	event.Get("/", eventController.GetEvents)
	event.Patch("/photo/:eventId", eventController.ConfirmNewPhoto)
	event.Post("/photo/:eventId", eventController.RequestNewPhoto)
	event.Patch("/:eventId", eventController.UpdateEvent)

	event.Post("/:eventId/subscriber", eventController.SubscribeIntoEvent)

	utils.Log.Info("EVENT routes successfully set up.")
}
