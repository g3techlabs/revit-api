package controller

import (
	"strconv"

	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/core/event/service"
	"github.com/g3techlabs/revit-api/response/generics"
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

func (c *EventController) CreateEvent(ctx *fiber.Ctx) error {
	var data input.CreateEventInput
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId := ctx.Locals("userId").(uint)

	response, err := c.eventService.CreateEvent(userId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *EventController) ConfirmNewPhoto(ctx *fiber.Ctx) error {
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

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.ConfirmNewPhoto(userId, eventId, &data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *EventController) GetEvents(ctx *fiber.Ctx) error {
	var data input.GetEventsFilters
	if err := ctx.QueryParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.eventService.GetEvents(userId, &data)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
