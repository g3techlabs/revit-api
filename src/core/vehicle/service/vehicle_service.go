package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	"github.com/g3techlabs/revit-api/src/core/vehicle/repository"
	"github.com/g3techlabs/revit-api/src/core/vehicle/response"
	"github.com/g3techlabs/revit-api/src/infra/storage"
	"github.com/g3techlabs/revit-api/src/validation"
)

type IVehicleService interface {
	CreateVehicle(userId uint, data *input.CreateVehicle) (*response.PresignedPhotoInfo, error)
	GetVehicles(userId uint, query *input.GetVehiclesParams) (*response.GetVehiclesResponse, error)
	GetVehicle(userId, vehicleId uint) (*response.GetVehicleResponse, error)
	ConfirmNewPhoto(userId, vehicleId uint, data *input.ConfirmNewPhoto) error
	UpdateVehicleInfo(vehicleId uint, data *input.UpdateVehicleInfo) error
	RequestPhotoUpsert(userId, vehicleId uint, data *input.RequestPhotoUpsert) (*response.PresignedPhotoInfo, error)
	DeleteVehicle(userId, vehicleId uint) error
	RemoveMainPhoto(userId, vehicleId uint) error
	RemovePhoto(userId, vehicleId, photoId uint) error
}

type VehicleService struct {
	vehicleRepo    repository.VehicleRepository
	validator      validation.IValidator
	storageService storage.StorageService
}

func NewVehicleService(validator validation.IValidator, vehicleRepo repository.VehicleRepository, storageService storage.StorageService) IVehicleService {
	return &VehicleService{vehicleRepo: vehicleRepo, validator: validator, storageService: storageService}
}
