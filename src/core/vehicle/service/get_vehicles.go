package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	"github.com/g3techlabs/revit-api/src/core/vehicle/response"
	"github.com/g3techlabs/revit-api/src/db/models"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (vs *VehicleService) GetVehicles(userId uint, query *input.GetVehiclesParams) (*[]response.Vehicle, error) {
	vehicles, err := vs.vehicleRepo.GetVehicles(userId, query.Page, query.Limit, query.Nickname)
	if err != nil {
		return nil, generics.InternalError()
	}

	vehiclesResponse := make([]response.Vehicle, 0, len(*vehicles))
	for _, vehicle := range *vehicles {
		mainPhotoUrl := buildMainPhotoUrl(vehicle.MainPhoto)
		photosResponse := mapPhotos(vehicle.Photos)

		vehiclesResponse = append(vehiclesResponse, response.Vehicle{
			ID:           vehicle.ID,
			Nickname:     vehicle.Nickname,
			Brand:        vehicle.Brand,
			Model:        vehicle.Model,
			Year:         vehicle.Year,
			Version:      vehicle.Version,
			MainPhotoUrl: mainPhotoUrl,
			CreatedAt:    vehicle.CreatedAt,
			Photos:       photosResponse,
		})
	}

	return &vehiclesResponse, nil
}

func buildMainPhotoUrl(mainPhoto *string) *string {
	if mainPhoto == nil {
		return nil
	}
	return utils.MountCloudFrontUrl(*mainPhoto)
}

func mapPhotos(photos []models.Photo) []response.Photo {
	result := make([]response.Photo, 0, len(photos))
	for _, photo := range photos {
		url := utils.MountCloudFrontUrl(photo.Reference)
		result = append(result, response.Photo{
			ID:  photo.ID,
			Url: *url,
		})
	}
	return result
}
