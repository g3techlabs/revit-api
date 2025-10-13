package service

import (
	"github.com/g3techlabs/revit-api/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/core/vehicle/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (vs *VehicleService) ConfirmNewMainPhoto(vehicleId uint, data *input.ConfirmNewMainPhoto) error {
	if err := vs.validator.Validate(data); err != nil {
		return err
	}

	if err := vs.storageService.DoesObjectExist(data.ObjectKey); err != nil {
		return err
	}

	if err := vs.vehicleRepo.UpdateMainPhoto(vehicleId, data.ObjectKey); err != nil {
		if err.Error() == "vehicle not found" {
			return errors.VehicleNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
