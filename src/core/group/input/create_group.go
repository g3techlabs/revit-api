package input

import (
	"github.com/g3techlabs/revit-api/src/db/models"
)

var publicVisibility uint = 1
var privateVisibility uint = 2

// CreateGroup representa os dados para criação de um novo grupo
// @Description Dados necessários para criar um novo grupo no sistema
type CreateGroup struct {
	// Nome do grupo
	Name string `json:"name" validate:"required" example:"Grupo de Ciclistas"`
	// Descrição do grupo
	Description string `json:"description" validate:"required" example:"Grupo para ciclistas da cidade"`
	// Tipo de conteúdo da foto principal (opcional: image/jpeg, image/png, image/webp)
	MainPhotoContentType *string `json:"mainPhotoContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp" example:"image/jpeg"`
	// Tipo de conteúdo do banner (opcional: image/jpeg, image/png, image/webp)
	BannerContentType *string `json:"bannerContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp" example:"image/png"`
	// Visibilidade do grupo (public ou private)
	Visibility string `json:"visibility" validate:"required,oneof=public private" example:"public"`
	// ID da cidade onde o grupo está localizado
	CityID uint `json:"cityId" validate:"required,number,gt=0" example:"1"`
}

func (i *CreateGroup) ToGroupModel() *models.Group {
	groupModel := new(models.Group)

	groupModel.Name = i.Name
	groupModel.Description = i.Description
	groupModel.CityID = i.CityID

	switch i.Visibility {
	case "public":
		groupModel.VisibilityID = publicVisibility
	case "private":
		groupModel.VisibilityID = privateVisibility
	}

	return groupModel
}
