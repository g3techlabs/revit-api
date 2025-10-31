package service

import (
	"github.com/g3techlabs/revit-api/src/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/src/core/vehicle/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (vs *VehicleService) ConfirmNewPhoto(userId, vehicleId uint, data *input.ConfirmNewPhoto) error {
	if err := vs.validator.Validate(data); err != nil {
		return err
	}

	if err := vs.storageService.DoesObjectExist(data.ObjectKey); err != nil {
		return err
	}

	switch data.PhotoType {
	case "main":
		if err := vs.vehicleRepo.UpdateMainPhoto(vehicleId, data.ObjectKey); err != nil {
			if err.Error() == "vehicle not found" {
				return errors.VehicleNotFound()
			}
			return generics.InternalError()
		}
	case "feed":
		vehicleAvailable, err := vs.vehicleRepo.IsVehicleAvailable(userId, vehicleId)
		if err != nil {
			return generics.InternalError()
		} else if !vehicleAvailable {
			return errors.VehicleNotFound()
		}

		if err := vs.vehicleRepo.InsertPhoto(vehicleId, data.ObjectKey); err != nil {
			return generics.InternalError()
		}
	}

	return nil
}
