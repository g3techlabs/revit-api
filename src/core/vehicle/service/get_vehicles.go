package service

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	"github.com/g3techlabs/revit-api/src/core/vehicle/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (vs *VehicleService) GetVehicles(userId uint, query *input.GetVehiclesParams) (*response.GetVehiclesResponse, error) {
	vehicles, totalCount, err := vs.vehicleRepo.GetVehicles(userId, query.Page, query.Limit, query.Nickname)
	if err != nil {
		return nil, generics.InternalError()
	}

	limit := 20
	if query.Limit > 0 {
		limit = int(query.Limit)
	}

	totalPages := uint(0)
	if totalCount > 0 && limit > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limit)))
	}

	currentPage := uint(1)
	if query.Page > 0 {
		currentPage = query.Page
	}

	vehiclesResponse := make([]response.SimpleVehicle, 0, len(*vehicles))
	for _, vehicle := range *vehicles {
		mainPhotoUrl := buildMainPhotoUrl(vehicle.MainPhoto)

		vehiclesResponse = append(vehiclesResponse, response.SimpleVehicle{
			ID:           vehicle.ID,
			Nickname:     vehicle.Nickname,
			Brand:        vehicle.Brand,
			Model:        vehicle.Model,
			Year:         vehicle.Year,
			Version:      vehicle.Version,
			MainPhotoUrl: mainPhotoUrl,
		})
	}

	return &response.GetVehiclesResponse{
		Vehicles:    vehiclesResponse,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}, nil
}

func buildMainPhotoUrl(mainPhoto *string) *string {
	if mainPhoto == nil {
		return nil
	}
	return utils.MountCloudFrontUrl(*mainPhoto)
}
