package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (vs *VehicleService) UpdateVehicleInfo(vehicleId uint, data *input.UpdateVehicleInfo) error {
	if err := vs.validator.Validate(data); err != nil {
		return err
	}

	if err := vs.vehicleRepo.UpdateVehicleInfo(vehicleId, data.ToVehicleModel()); err != nil {
		if err.Error() == "vehicle not found" {
			return errors.VehicleNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
