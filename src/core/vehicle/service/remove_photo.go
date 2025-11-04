package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (vs *VehicleService) RemovePhoto(userId, vehicleId, photoId uint) error {
	if err := vs.markFeedPhotoAsRemoved(userId, vehicleId, photoId); err != nil {
		return err
	}

	return nil
}

func (vs *VehicleService) markFeedPhotoAsRemoved(userId, vehicleId, photoId uint) error {
	if err := vs.vehicleRepo.MarkPhotoAsRemoved(userId, vehicleId, photoId); err != nil {
		if err.Error() == "photo not found" {
			return errors.PhotoNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
