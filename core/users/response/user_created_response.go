package response

import "time"

type UserCreatedResponse struct {
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Nickname   string     `json:"nickname"`
	ProfilePic *string    `json:"profilePic"`
	Birthdate  *time.Time `json:"birthdate"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}
