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

func (c *EventController) UpdateEvent(ctx *fiber.Ctx) error {
	var data input.UpdateEventInput
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
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

	if err := c.eventService.UpdateEvent(userId, eventId, &data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *EventController) RequestNewPhoto(ctx *fiber.Ctx) error {
	var data input.RequestNewPhotoInput
	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
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

	response, err := c.eventService.RequestNewPhoto(userId, eventId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *EventController) SubscribeIntoEvent(ctx *fiber.Ctx) error {
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

	if err := c.eventService.SubscribeToEvent(userId, eventId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *EventController) RevokeEventSubscription(ctx *fiber.Ctx) error {
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

	if err := c.eventService.RevokeEventSubscription(userId, eventId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *EventController) InviteUserToEvent(ctx *fiber.Ctx) error {
	eventParam := ctx.Params("eventId")
	eventIdUint64, err := strconv.ParseUint(eventParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID",
		})
	}
	eventId := uint(eventIdUint64)

	invitedParam := ctx.Params("invitedId")
	invitedIdUint64, err := strconv.ParseUint(invitedParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid invite target ID",
		})
	}
	invitedId := uint(invitedIdUint64)

	eventAdminId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.eventService.InviteUserToEvent(eventAdminId, eventId, invitedId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *EventController) GetPendingInvites(ctx *fiber.Ctx) error {
	var filters input.GetPendingInvitesFilters

	if err := ctx.QueryParser(&filters); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.eventService.GetPendingInvites(userId, &filters)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *EventController) AnswerPendingInvite(ctx *fiber.Ctx) error {
	var answer input.PendingInviteAnswer

	if err := ctx.BodyParser(&answer); err != nil {
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

	if err := c.eventService.AnswerPendingInvite(userId, eventId, &answer); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
