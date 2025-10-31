package response

import "time"

type Vehicle struct {
	ID           uint      `json:"id"`
	Nickname     string    `json:"nickname"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	Year         uint      `json:"year"`
	Version      *string   `json:"version"`
	MainPhotoUrl *string   `json:"mainPhotoUrl"`
	CreatedAt    time.Time `json:"createdAt"`
	Photos       []Photo
}

type Photo struct {
	ID  uint   `json:"id"`
	Url string `json:"reference"`
}
