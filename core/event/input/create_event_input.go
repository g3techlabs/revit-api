package input

import (
	"time"

	"github.com/g3techlabs/revit-api/core/event/entities"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

type CreateEventInput struct {
	Name             string               `json:"name" validate:"required"`
	Description      string               `json:"description" validate:"required"`
	Date             string               `json:"date" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Location         entities.Coordinates `json:"location" validate:"required"`
	PhotoContentType *string              `json:"photoContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp"`
	City             string               `json:"city" validate:"required"`
	Visibility       string               `json:"visibility" validate:"required,oneof=public private"`
	GroupID          *uint                `json:"groupId" validate:"omitempty,number,gt=0"`
}

func (i *CreateEventInput) ToEventModel() *models.Event {
	dateTime, err := time.Parse("2006-01-02T15:04:05Z07:00", i.Date)
	if err != nil {
		return nil
	}

	var newVisibility uint
	switch i.Visibility {
	case "public":
		newVisibility = 1
	case "private":
		newVisibility = 2
	}

	return &models.Event{
		Name:         i.Name,
		Description:  i.Description,
		Date:         dateTime,
		City:         i.City,
		VisibilityID: newVisibility,
		GroupID:      i.GroupID,
		Location:     gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", i.Location.Longitude, i.Location.Latitude),
	}
}
