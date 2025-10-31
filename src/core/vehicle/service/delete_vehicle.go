package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (vr *VehicleService) DeleteVehicle(userId, vehicleId uint) error {
	if err := vr.vehicleRepo.MarkVehicleAsRemoved(userId, vehicleId); err != nil {
		if err.Error() == "vehicle not found" {
			return errors.VehicleNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
