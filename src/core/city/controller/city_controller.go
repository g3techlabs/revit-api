package controller

import (
	"github.com/g3techlabs/revit-api/src/core/city/input"
	_ "github.com/g3techlabs/revit-api/src/core/city/response"
	"github.com/g3techlabs/revit-api/src/core/city/service"
	"github.com/gofiber/fiber/v2"
)

type CityController struct {
	cityService service.ICityService
}

func NewCityController(cityService service.ICityService) *CityController {
	return &CityController{cityService: cityService}
}

func (c *CityController) GetCities(ctx *fiber.Ctx) error {
	var query input.GetCitiesFilters

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	response, err := c.cityService.GetCities(&query)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *CityController) GetNearbyCities(ctx *fiber.Ctx) error {
	var query input.GetNearbyCitiesFilters

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	response, err := c.cityService.GetNearbyCities(&query)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
