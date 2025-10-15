package router

import (
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/token"
	userRepo "github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/core/vehicle/controller"
	"github.com/g3techlabs/revit-api/core/vehicle/service"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
)

func VehicleRoutes(router fiber.Router, vehicleService service.IVehicleService, ur userRepo.UserRepository, ts token.ITokenService) {
	utils.Log.Info("Setting up VEHICLE routes...")

	vehicleController := controller.NewVehicleController(vehicleService)

	vehicle := router.Group("/vehicle", middleware.Auth(ur, ts))

	vehicle.Post("/", vehicleController.CreateVehicle)
	vehicle.Get("/", vehicleController.GetVehicles)
	vehicle.Patch("/photo/:vehicleId", vehicleController.ConfirmNewPhoto)
	vehicle.Post("/photo/:vehicleId", vehicleController.RequestPhotoUpsert)
	vehicle.Patch("/:vehicleId", vehicleController.UpdateVehicleInfo)
	vehicle.Delete("/:vehicleId/photo/:photoId", vehicleController.RemovePhoto)
	vehicle.Delete("/main-photo/:vehicleId", vehicleController.RemoveMainPhoto)

	utils.Log.Info("VEHICLE routes successfully set up.")
}
