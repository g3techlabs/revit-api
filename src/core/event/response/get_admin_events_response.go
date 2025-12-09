package response

import "time"

type AdminEvent struct {
	ID   uint      `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type GetAdminEventsResponse struct {
	Events      []AdminEvent `json:"events"`
	TotalPages  uint         `json:"totalPages"`
	CurrentPage uint         `json:"currentPage"`
}
