package controller

import (
	"github.com/g3techlabs/revit-api/core/city/input"
	"github.com/g3techlabs/revit-api/core/city/service"
	"github.com/gofiber/fiber/v2"
)

type CityController struct {
	cityService service.ICityService
}

func NewCityController(cityService service.ICityService) *CityController {
	return &CityController{cityService: cityService}
}

func (c *CityController) GetCities(ctx *fiber.Ctx) error {
	var query input.GetCitiesQuery

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	response, err := c.cityService.GetCities(&query)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
