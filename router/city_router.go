package router

import (
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/city/controller"
	"github.com/g3techlabs/revit-api/core/city/service"
	"github.com/gofiber/fiber/v2"
)

func CityRoutes(router fiber.Router, cityService service.ICityService, m *middleware.AuthMiddleware) {
	controller := controller.NewCityController(cityService)

	city := router.Group("/city", m.Auth())
	city.Get("/", controller.GetCities)
}
