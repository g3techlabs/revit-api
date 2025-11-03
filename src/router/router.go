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
	gls "github.com/g3techlabs/revit-api/src/core/geolocation"
	glr "github.com/g3techlabs/revit-api/src/core/geolocation/repository"
	gr "github.com/g3techlabs/revit-api/src/core/group/repository"
	gs "github.com/g3techlabs/revit-api/src/core/group/service"
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

func SetupRoutes(app *fiber.App) {
	utils.Log.Info("Setting up routes...")

	validator := validation.NewValidator()
	storageClient := config.NewS3Client()
	hub := websocket.NewHub()
	redisClient := config.NewRedisClient()

	userRepo := repository.NewUserRepository()
	vehicleRepo := vr.NewVehicleRepository()
	groupRepository := gr.NewGroupRepository()
	eventRepository := er.NewEventRepository()
	cityRepository := cr.NewCityRepository()
	geoLocationRepo := glr.NewGeoLocationRepository(context.TODO(), redisClient)

	storageService := storage.NewS3Service(storageClient, config.NewPresignClient(storageClient), context.Background())
	tokenService := token.NewTokenService()
	emailService := mail.NewEmailService()

	authMiddleware := middleware.NewAuthMiddleware(userRepo, tokenService)

	authService := services.NewAuthService(validator, userRepo, emailService, tokenService)
	userService := service.NewUserService(validator, userRepo, storageService)
	vehicleService := vs.NewVehicleService(validator, vehicleRepo, storageService)
	groupService := gs.NewGroupService(groupRepository, validator, storageService)
	eventService := es.NewEventService(validator, eventRepository, storageService)
	cityService := cs.NewCityService(validator, cityRepository)
	geoLocationService := gls.NewGeoLocationService(validator, geoLocationRepo, hub)

	api := app.Group("/api")
	AuthRoutes(api, authService)
	UserRoutes(api, userService, authMiddleware)
	VehicleRoutes(api, vehicleService, authMiddleware)
	GroupRoutes(api, groupService, authMiddleware)
	EventRoutes(api, eventService, authMiddleware)
	CityRoutes(api, cityService, authMiddleware)
	WebSocketRoute(api, hub, geoLocationService, authMiddleware)

	utils.Log.Info("Routes successfully set up.")
}
