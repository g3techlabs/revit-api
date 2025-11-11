package router

import (
	"context"

	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/core/auth/middleware"
	"github.com/g3techlabs/revit-api/src/core/auth/services"
	cr "github.com/g3techlabs/revit-api/src/core/city/repository"
	cs "github.com/g3techlabs/revit-api/src/core/city/service"
	er "github.com/g3techlabs/revit-api/src/core/event/repository"
	es "github.com/g3techlabs/revit-api/src/core/event/service"
	glr "github.com/g3techlabs/revit-api/src/core/geolocation/repository"
	gls "github.com/g3techlabs/revit-api/src/core/geolocation/service"
	gr "github.com/g3techlabs/revit-api/src/core/group/repository"
	gs "github.com/g3techlabs/revit-api/src/core/group/service"
	rr "github.com/g3techlabs/revit-api/src/core/route/repository"
	rs "github.com/g3techlabs/revit-api/src/core/route/service"
	"github.com/g3techlabs/revit-api/src/core/users/repository"
	"github.com/g3techlabs/revit-api/src/core/users/service"
	vr "github.com/g3techlabs/revit-api/src/core/vehicle/repository"
	vs "github.com/g3techlabs/revit-api/src/core/vehicle/service"
	"github.com/g3techlabs/revit-api/src/infra/mail"
	"github.com/g3techlabs/revit-api/src/infra/storage"
	"github.com/g3techlabs/revit-api/src/infra/token"
	"github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/g3techlabs/revit-api/src/validation"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, logger utils.ILogger) {
	logger.Info("Setting up routes...")

	validator := validation.NewValidator()
	storageClient := config.NewS3Client()
	redisClient := config.NewRedisClient()
	hub := websocket.NewHub(logger)
	go hub.Run()

	userRepo := repository.NewUserRepository()
	vehicleRepo := vr.NewVehicleRepository()
	groupRepository := gr.NewGroupRepository()
	eventRepository := er.NewEventRepository()
	cityRepository := cr.NewCityRepository()
	geoLocationRepo := glr.NewGeoLocationRepository(redisClient)
	routeRepo := rr.NewRouteRepository()

	storageService := storage.NewS3Service(storageClient, config.NewPresignClient(storageClient), context.Background(), logger)
	tokenService := token.NewTokenService()
	emailService := mail.NewEmailService()

	authMiddleware := middleware.NewAuthMiddleware(userRepo, tokenService)

	authService := services.NewAuthService(validator, userRepo, emailService, tokenService)
	userService := service.NewUserService(validator, userRepo, storageService)
	vehicleService := vs.NewVehicleService(validator, vehicleRepo, storageService)
	groupService := gs.NewGroupService(groupRepository, validator, storageService)
	eventService := es.NewEventService(validator, eventRepository, storageService)
	cityService := cs.NewCityService(validator, cityRepository)
	geoLocationService := gls.NewGeoLocationService(validator, geoLocationRepo, hub, logger)
	routeService := rs.NewRouteService(validator, geoLocationService, routeRepo, hub)

	api := app.Group("/api")
	AuthRoutes(api, authService, logger)
	UserRoutes(api, userService, authMiddleware, logger)
	VehicleRoutes(api, vehicleService, authMiddleware, logger)
	GroupRoutes(api, groupService, authMiddleware, logger)
	EventRoutes(api, eventService, authMiddleware, logger)
	CityRoutes(api, cityService, authMiddleware, logger)
	WebSocketRoute(api, hub, routeService, geoLocationService, authMiddleware, logger)
	RouteRoutes(api, routeService, authMiddleware, logger)

	logger.Info("Routes successfully set up.")
}
