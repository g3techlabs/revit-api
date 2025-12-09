package input

import "github.com/g3techlabs/revit-api/src/db/models"

type UpdateVehicleInfo struct {
	Nickname *string `json:"nickname" validate:"omitempty"`
	Brand    *string `json:"brand" validate:"omitempty"`
	Model    *string `json:"model" validate:"omitempty"`
	Year     *uint   `json:"year" validate:"omitempty,number,gt=0"`
	Version  *string `json:"version" validate:"omitempty"`
}

func (u *UpdateVehicleInfo) ToVehicleModel() *models.Vehicle {
	model := new(models.Vehicle)

	if u.Nickname != nil {
		model.Nickname = *u.Nickname
	}
	if u.Brand != nil {
		model.Brand = *u.Brand
	}
	if u.Model != nil {
		model.Model = *u.Model
	}
	if u.Year != nil {
		model.Year = *u.Year
	}
	if u.Version != nil {
		model.Version = u.Version
	}

	return model
}
