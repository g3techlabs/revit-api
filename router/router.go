package router

import (
	"context"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/auth/services"
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

	storageService := storage.NewS3Service(storageClient, config.NewPresignClient(storageClient), context.Background())
	tokenService := token.NewTokenService()
	emailService := mail.NewEmailService()

	authService := services.NewAuthService(validator, userRepo, emailService, tokenService)
	userService := service.NewUserService(validator, userRepo, storageService)
	vehicleService := vs.NewVehicleService(validator, vehicleRepo, storageService)

	api := app.Group("/api")
	AuthRoutes(api, authService)
	UserRoutes(api, userService, userRepo, tokenService)
	VehicleRoutes(api, vehicleService, userRepo, tokenService)

	utils.Log.Info("Routes successfully set up.")
}
