package response

import "time"

type GetUserResponse struct {
	ID            uint       `json:"id"`
	Name          string     `json:"name"`
	Nickname      string     `json:"nickname"`
	ProfilePicUrl *string    `json:"ProfilePicUrl"`
	Since         *time.Time `json:"since"`
}
