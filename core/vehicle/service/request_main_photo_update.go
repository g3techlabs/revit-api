package service

import (
	"github.com/g3techlabs/revit-api/core/vehicle/input"
	"github.com/g3techlabs/revit-api/core/vehicle/response"
)

func (vs *VehicleService) RequestMainPhotoUpdate(userId, vehicleId uint, data *input.RequestMainPhotoUpdate) (*response.PresignedMainPhotoInfo, error) {
	if err := vs.validator.Validate(data); err != nil {
		return nil, err
	}

	response, err := vs.buildPresignedMainPhotoResponse(userId, vehicleId, &data.ContentType)
	if err != nil {
		return nil, err
	}

	return response, nil
}
