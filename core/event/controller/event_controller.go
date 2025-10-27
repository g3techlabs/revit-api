package controller

import (
	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/core/event/service"
	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	eventService service.IEventService
}

func NewEventController(eventService service.IEventService) *EventController {
	return &EventController{
		eventService: eventService,
	}
}

func (ec *EventController) CreateEvent(ctx *fiber.Ctx) error {
	var data input.CreateEventInput
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId := ctx.Locals("userId").(uint)

	if _, err := ec.eventService.CreateEvent(userId, &data); err != nil {
		return err
	}

	return nil
}
