package service

import (
	"github.com/g3techlabs/revit-api/core/storage"
	"github.com/g3techlabs/revit-api/core/vehicle/input"
	"github.com/g3techlabs/revit-api/core/vehicle/repository"
	"github.com/g3techlabs/revit-api/core/vehicle/response"
	"github.com/g3techlabs/revit-api/validation"
)

type IVehicleService interface {
	CreateVehicle(userId uint, data *input.CreateVehicle) (*response.VehicleCreated, error)
	ConfirmNewMainPhoto(vehicleId uint, data *input.ConfirmNewMainPhoto) error
}

type VehicleService struct {
	vehicleRepo    repository.VehicleRepository
	validator      validation.IValidator
	storageService storage.StorageService
}

func NewVehicleService(validator validation.IValidator, vehicleRepo repository.VehicleRepository, storageService storage.StorageService) IVehicleService {
	return &VehicleService{vehicleRepo: vehicleRepo, validator: validator, storageService: storageService}
}
