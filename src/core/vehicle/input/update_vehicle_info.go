package input

import "github.com/g3techlabs/revit-api/src/db/models"

// UpdateVehicleInfo representa os dados para atualização de um veículo
// @Description Dados opcionais para atualizar informações de um veículo existente
type UpdateVehicleInfo struct {
	// Apelido do veículo (opcional)
	Nickname *string `json:"nickname" validate:"omitempty" example:"Minha Moto Atualizada"`
	// Marca do veículo (opcional)
	Brand *string `json:"brand" validate:"omitempty" example:"Yamaha"`
	// Modelo do veículo (opcional)
	Model *string `json:"model" validate:"omitempty" example:"MT-07"`
	// Ano do veículo (opcional, deve ser maior que 0)
	Year *uint `json:"year" validate:"omitempty,number,gt=0" example:"2021"`
	// Versão do veículo (opcional)
	Version *string `json:"version" validate:"omitempty" example:"ABS"`
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
