package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	"github.com/g3techlabs/revit-api/src/core/vehicle/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (vs *VehicleService) RequestPhotoUpsert(userId, vehicleId uint, data *input.RequestPhotoUpsert) (*response.PresignedPhotoInfo, error) {
	if err := vs.validator.Validate(data); err != nil {
		return nil, err
	}

	vehicleAvailable, err := vs.vehicleRepo.IsVehicleAvailable(userId, vehicleId)
	if err != nil {
		return nil, generics.InternalError()
	} else if !vehicleAvailable {
		return nil, errors.VehicleNotFound()
	}

	response, err := vs.buildPresignedPhotoResponse(userId, vehicleId, &data.ContentType, data.PhotoType)
	if err != nil {
		return nil, err
	}

	return response, nil
}
