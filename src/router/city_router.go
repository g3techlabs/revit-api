package router

import (
	"github.com/g3techlabs/revit-api/src/core/auth/middleware"
	"github.com/g3techlabs/revit-api/src/core/city/controller"
	"github.com/g3techlabs/revit-api/src/core/city/service"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func CityRoutes(router fiber.Router, cityService service.ICityService, m *middleware.AuthMiddleware, logger utils.ILogger) {
	logger.Info("Setting up CITY routes...")

	controller := controller.NewCityController(cityService)

	city := router.Group("/city", m.Auth())
	city.Get("/", controller.GetCities)
	city.Get("/nearby", controller.GetNearbyCities)

	logger.Info("CITY routes successfully set up.")
}
