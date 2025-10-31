package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (vs *VehicleService) RemoveMainPhoto(userId, vehicleId uint) error {
	if err := vs.markMainPhotoAsRemoved(userId, vehicleId); err != nil {
		return err
	}

	return nil
}

func (vs *VehicleService) markMainPhotoAsRemoved(userId, vehicleId uint) error {
	if err := vs.vehicleRepo.DeleteMainPhoto(userId, vehicleId); err != nil {
		if err.Error() == "vehicle not found" {
			return errors.VehicleNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
