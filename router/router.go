package router

import (
	"context"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/auth/services"
	gr "github.com/g3techlabs/revit-api/core/group/repository"
	gs "github.com/g3techlabs/revit-api/core/group/service"
	"github.com/g3techlabs/revit-api/core/mail"
	"github.com/g3techlabs/revit-api/core/storage"
	"github.com/g3techlabs/revit-api/core/token"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/core/users/service"
	vr "github.com/g3techlabs/revit-api/core/vehicle/repository"
	vs "github.com/g3techlabs/revit-api/core/vehicle/service"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/g3techlabs/revit-api/validation"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	utils.Log.Info("Setting up routes...")

	validator := validation.NewValidator()
	storageClient := config.NewS3Client()

	userRepo := repository.NewUserRepository()
	vehicleRepo := vr.NewVehicleRepository()
	groupRepository := gr.NewGroupRepository()

	storageService := storage.NewS3Service(storageClient, config.NewPresignClient(storageClient), context.Background())
	tokenService := token.NewTokenService()
	emailService := mail.NewEmailService()

	authMiddleware := middleware.NewAuthMiddleware(userRepo, tokenService)

	authService := services.NewAuthService(validator, userRepo, emailService, tokenService)
	userService := service.NewUserService(validator, userRepo, storageService)
	vehicleService := vs.NewVehicleService(validator, vehicleRepo, storageService)
	groupService := gs.NewGroupService(groupRepository, validator, storageService)

	api := app.Group("/api")
	AuthRoutes(api, authService)
	UserRoutes(api, userService, authMiddleware)
	VehicleRoutes(api, vehicleService, authMiddleware)
	GroupRoutes(api, groupService, authMiddleware)

	utils.Log.Info("Routes successfully set up.")
}
