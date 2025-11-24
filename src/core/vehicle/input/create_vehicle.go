package input

import (
	"github.com/g3techlabs/revit-api/src/db/models"
)

// CreateVehicle representa os dados para criação de um novo veículo
// @Description Dados necessários para criar um novo veículo para o usuário autenticado
type CreateVehicle struct {
	// Apelido do veículo
	Nickname string `json:"nickname" validate:"required" example:"Minha Moto"`
	// Marca do veículo
	Brand string `json:"brand" validate:"required" example:"Honda"`
	// Modelo do veículo
	Model string `json:"model" validate:"required" example:"CB 600F"`
	// Ano do veículo (deve ser maior que 0)
	Year uint `json:"year" validate:"required,number,gt=0" example:"2020"`
	// Versão do veículo (opcional)
	Version *string `json:"version" validate:"omitempty" example:"Hornet"`
	// Tipo de conteúdo da foto principal (opcional: image/jpeg, image/png, image/webp)
	MainPhotoContentTye *string `json:"mainPhotoContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp" example:"image/jpeg"`
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
