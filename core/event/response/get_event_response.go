package response

import (
	"time"

	"gorm.io/datatypes"
)

type GetEventResponse struct {
	ID               uint            `json:"id"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	Date             time.Time       `json:"date"`
	Visibility       string          `json:"visibility"`
	Photo            *string         `json:"photo"`
	SubscribersCount *uint           `json:"subscribersCount"`
	Address          datatypes.JSON  `json:"address"`
	Coordinates      datatypes.JSON  `json:"coordinates"`
	Group            *datatypes.JSON `json:"group"`
	MemberRole       string          `json:"memberRole"`
}
