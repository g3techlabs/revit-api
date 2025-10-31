package input

import "github.com/g3techlabs/revit-api/src/core/event/entities"

type UpdateEventInput struct {
	Name        *string               `json:"name" validate:"omitempty"`
	Description *string               `json:"description" validate:"omitempty"`
	Date        *string               `json:"date" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	Location    *entities.Coordinates `json:"location" validate:"required_with=CityID,omitempty"`
	CityID      *uint                 `json:"cityId" validate:"required_with=Location,omitempty,number,gt=0"`
	GroupID     *uint                 `json:"groupId" validate:"omitempty,number,gt=0"`
	Visibility  *string               `json:"visibility" validate:"omitempty,oneof=public private"`
}
