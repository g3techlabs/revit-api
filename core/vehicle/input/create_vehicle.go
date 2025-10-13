package input

import (
	"github.com/g3techlabs/revit-api/db/models"
)

type CreateVehicle struct {
	Nickname            string  `json:"nickname" validate:"required"`
	Brand               string  `json:"brand" validate:"required"`
	Model               string  `json:"model" validate:"required"`
	Year                uint    `json:"year" validate:"required,number,gt=0"`
	Version             *string `json:"version" validate:"omitempty"`
	MainPhotoContentTye *string `json:"mainPhotoContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp"`
}

func (i CreateVehicle) ToVehicleModel(userId uint) *models.Vehicle {
	return &models.Vehicle{
		Nickname: i.Nickname,
		Brand:    i.Brand,
		Model:    i.Model,
		Year:     i.Year,
		Version:  i.Version,
		UserID:   userId,
	}
}
