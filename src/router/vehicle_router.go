package router

import (
	"github.com/g3techlabs/revit-api/src/core/auth/middleware"
	"github.com/g3techlabs/revit-api/src/core/vehicle/controller"
	"github.com/g3techlabs/revit-api/src/core/vehicle/service"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func VehicleRoutes(router fiber.Router, vehicleService service.IVehicleService, middleware *middleware.AuthMiddleware, logger utils.ILogger) {
	logger.Info("Setting up VEHICLE routes...")

	vehicleController := controller.NewVehicleController(vehicleService)

	vehicle := router.Group("/vehicle", middleware.Auth())

	vehicle.Post("/", vehicleController.CreateVehicle)
	vehicle.Get("/", vehicleController.GetVehicles)
	vehicle.Patch("/photo/:vehicleId", vehicleController.ConfirmNewPhoto)
	vehicle.Post("/photo/:vehicleId", vehicleController.RequestPhotoUpsert)
	vehicle.Patch("/:vehicleId", vehicleController.UpdateVehicleInfo)
	vehicle.Delete("/:vehicleId/photo/:photoId", vehicleController.RemovePhoto)
	vehicle.Delete("/main-photo/:vehicleId", vehicleController.RemoveMainPhoto)
	vehicle.Delete("/:vehicleId", vehicleController.DeleteVehicle)

	logger.Info("VEHICLE routes successfully set up.")
}
