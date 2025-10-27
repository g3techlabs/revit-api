package controller

import (
	"strconv"

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

	response, err := ec.eventService.CreateEvent(userId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (ec *EventController) ConfirmNewPhoto(ctx *fiber.Ctx) error {
	var data input.ConfirmNewPhoto
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	userId := ctx.Locals("userId").(uint)

	if err := ec.eventService.ConfirmNewPhoto(userId, eventId, &data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
