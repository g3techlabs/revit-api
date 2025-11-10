package controller

import (
	"strconv"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/core/route/service"
	"github.com/gofiber/fiber/v2"
)

type RouteController struct {
	routeService service.IRouteService
}

func NewRouteController(routeService service.IRouteService) *RouteController {
	return &RouteController{
		routeService: routeService,
	}
}

func (c *RouteController) CreateRoute(ctx *fiber.Ctx) error {
	var data input.CreateRouteInput

	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	response, err := c.routeService.CreateRoute(userId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *RouteController) InviteUsers(ctx *fiber.Ctx) error {
	var data input.UsersToInviteInput

	if err := ctx.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	routeParam := ctx.Params("routeId")
	routeIdUint64, err := strconv.ParseUint(routeParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}
	routeId := uint(routeIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	if err := c.routeService.InviteUsers(userId, routeId, &data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *RouteController) GetOnlineFriendsToInvite(ctx *fiber.Ctx) error {
	var query input.GetOnlineFriendsToInviteQuery

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	response, err := c.routeService.GetOnlineFriendsToInvite(userId, &query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *RouteController) GetNearbyUsersToRouteInvite(ctx *fiber.Ctx) error {
	var data input.GetNearbyUsersToInviteQuery

	if err := ctx.QueryParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	response, err := c.routeService.GetNearbyUsersToInvite(userId, &data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *RouteController) AcceptRouteInvite(ctx *fiber.Ctx) error {
	var coordinates geoinput.Coordinates

	if err := ctx.BodyParser(&coordinates); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	routeParam := ctx.Params("routeId")
	routeIdUint64, err := strconv.ParseUint(routeParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}
	routeId := uint(routeIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	if err := c.routeService.AcceptRouteInvite(userId, routeId, &coordinates); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
