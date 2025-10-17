package input

import (
	"github.com/g3techlabs/revit-api/core/group/repository"
	"github.com/g3techlabs/revit-api/db/models"
)

type CreateGroup struct {
	Name                 string  `json:"name" validate:"required"`
	Description          string  `json:"description" validate:"required"`
	MainPhotoContentType *string `json:"mainPhotoContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp"`
	BannerContentType    *string `json:"bannerContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp"`
	Visibility           string  `json:"visibility" validate:"required,oneof=public private"`
	CityID               uint    `json:"cityId" validate:"required,number,gt=0"`
}

func (i *CreateGroup) ToGroupModel() *models.Group {
	groupModel := new(models.Group)

	groupModel.Name = i.Name
	groupModel.Description = i.Description
	groupModel.CityID = i.CityID

	switch i.Visibility {
	case "public":
		groupModel.VisibilityID = repository.PublicVisibility
	case "private":
		groupModel.VisibilityID = repository.PrivateVisibility
	}

	return groupModel
}
