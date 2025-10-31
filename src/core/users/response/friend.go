package response

import "time"

type Friend struct {
	ID            uint       `json:"id"`
	Name          string     `json:"name"`
	Nickname      string     `json:"nickname"`
	ProfilePicUrl *string    `json:"ProfilePicUrl"`
	Since         time.Time  `json:"since"`
	FriendsSince  *time.Time `json:"friendsSince"`
}
