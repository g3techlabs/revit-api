package response

import "time"

type UserCreatedResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Nickname   string    `json:"nickname"`
	ProfilePic *string   `json:"profilePic"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
