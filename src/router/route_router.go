package router

import (
	"github.com/g3techlabs/revit-api/src/core/auth/middleware"
	"github.com/g3techlabs/revit-api/src/core/route/controller"
	"github.com/g3techlabs/revit-api/src/core/route/service"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func RouteRoutes(router fiber.Router, routeService service.IRouteService, m *middleware.AuthMiddleware, logger utils.ILogger) {
	logger.Info("Setting up ROUTE endpoints")

	controller := controller.NewRouteController(routeService)

	route := router.Group("/route", m.Auth())

	route.Post("/", controller.CreateRoute)
	route.Get("/friends", controller.GetOnlineFriendsToInvite)
	route.Get("/nearby", controller.GetNearbyUsersToRouteInvite)
	route.Post("/:routeId/invite", controller.InviteNearbyUsers)
	route.Post("/:routeId/invite/action", controller.AcceptRouteInvite)

	logger.Info("ROUTE endpoints successfully set up")
}
