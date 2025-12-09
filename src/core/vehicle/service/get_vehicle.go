package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/src/core/vehicle/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (vs *VehicleService) GetVehicle(userId, vehicleId uint) (*response.GetVehicleResponse, error) {
	vehicle, err := vs.vehicleRepo.GetVehicle(userId, vehicleId)
	if err != nil {
		return nil, generics.InternalError()
	} else if vehicle == nil {
		return nil, errors.VehicleNotFound()
	}

	mainPhoto := buildMainPhotoUrl(vehicle.MainPhoto)

	photos := make([]response.VehiclePhoto, 0, len(vehicle.Photos))
	for _, photo := range vehicle.Photos {
		photoUrl := utils.MountCloudFrontUrl(photo.Reference)
		photos = append(photos, response.VehiclePhoto{
			ID:        photo.ID,
			Reference: *photoUrl,
			VehicleID: photo.VehicleID,
		})
	}

	return &response.GetVehicleResponse{
		Nickname:  vehicle.Nickname,
		Brand:     vehicle.Brand,
		Model:     vehicle.Model,
		Year:      vehicle.Year,
		Version:   vehicle.Version,
		MainPhoto: mainPhoto,
		CreatedAt: vehicle.CreatedAt,
		Photos:    photos,
	}, nil
}
