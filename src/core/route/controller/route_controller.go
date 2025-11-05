package controller

import (
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
