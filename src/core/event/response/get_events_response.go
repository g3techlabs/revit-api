package response

import (
	"time"

	"gorm.io/datatypes"
)

type SimpleEvent struct {
	ID                 uint            `json:"id"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	Date               time.Time       `json:"date"`
	Visibility         string          `json:"visibility"`
	Photo              *string         `json:"photo"`
	SubscribersCount   *uint           `json:"subscribersCount"`
	Address            datatypes.JSON  `json:"address"`
	Coordinates        datatypes.JSON  `json:"coordinates"`
	Group              *datatypes.JSON `json:"group"`
	MemberRole         *string         `json:"memberRole"`
	FriendsSubscribers datatypes.JSON  `json:"friendsSubscribers"`
	Host               *datatypes.JSON `json:"host"`
}

type GetEventsResponse struct {
	Events      []SimpleEvent `json:"events"`
	CurrentPage uint          `json:"currentPage"`
	TotalPages  uint          `json:"totalPages"`
}
