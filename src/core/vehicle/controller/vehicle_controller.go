package controller

import (
	"strconv"

	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	"github.com/g3techlabs/revit-api/src/core/vehicle/service"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/gofiber/fiber/v2"
)

type VehicleController struct {
	vehicleService service.IVehicleService
}

func NewVehicleController(vehicleService service.IVehicleService) *VehicleController {
	return &VehicleController{vehicleService: vehicleService}
}

func (c *VehicleController) CreateVehicle(ctx *fiber.Ctx) error {
	input := new(input.CreateVehicle)

	if err := ctx.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request: "+err.Error())
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.vehicleService.CreateVehicle(userId, input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *VehicleController) GetVehicles(ctx *fiber.Ctx) error {
	query := new(input.GetVehiclesParams)

	if err := ctx.QueryParser(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	vehicles, err := c.vehicleService.GetVehicles(userId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(vehicles)
}

func (c *VehicleController) UpdateVehicleInfo(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	input := new(input.UpdateVehicleInfo)
	if err := ctx.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := c.vehicleService.UpdateVehicleInfo(vehicleId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *VehicleController) RequestPhotoUpsert(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	input := new(input.RequestPhotoUpsert)
	if err := ctx.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.vehicleService.RequestPhotoUpsert(userId, vehicleId, input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *VehicleController) ConfirmNewPhoto(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	input := new(input.ConfirmNewPhoto)
	if err := ctx.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.vehicleService.ConfirmNewPhoto(userId, vehicleId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *VehicleController) DeleteVehicle(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.vehicleService.DeleteVehicle(userId, vehicleId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *VehicleController) RemoveMainPhoto(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.vehicleService.RemoveMainPhoto(userId, vehicleId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *VehicleController) RemovePhoto(ctx *fiber.Ctx) error {
	vehicleParam := ctx.Params("vehicleId")

	vehicleIdUint64, err := strconv.ParseUint(vehicleParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle ID",
		})
	}
	vehicleId := uint(vehicleIdUint64)

	photoParam := ctx.Params("photoId")

	photoIdUint64, err := strconv.ParseUint(photoParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid photo ID",
		})
	}
	photoId := uint(photoIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.vehicleService.RemovePhoto(userId, vehicleId, photoId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
