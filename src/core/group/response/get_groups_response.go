package response

import (
	"time"

	"gorm.io/datatypes"
)

type GetGroupsResponse struct {
	ID             uint           `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	MainPhoto      *string        `json:"mainPhoto"`
	Banner         *string        `json:"banner"`
	CreatedAt      time.Time      `json:"createdAt"`
	Visibility     string         `json:"visibility"`
	City           string         `json:"city"`
	State          string         `json:"state"`
	MemberType     *string        `json:"memberType"`
	FriendsInGroup datatypes.JSON `json:"friendsInGroup"`
}
