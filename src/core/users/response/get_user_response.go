package response

import "time"

type GetUserResponse struct {
	Name          string        `json:"name" example:"João Silva"`
	Nickname      string        `json:"nickname" example:"joaosilva"`
	ProfilePicUrl *string       `json:"profilePicUrl" example:"https://example.com/users/42/profile.jpg"`
	Since         time.Time     `json:"since" example:"2024-01-15T10:30:00Z"`
	IsFriend      bool          `json:"isFriend" example:"true"`
	Vehicles      []UserVehicle `json:"vehicles"`
	Groups        []UserGroup   `json:"groups"`
	Events        []UserEvent   `json:"events"`
}

type UserVehicle struct {
	ID       uint           `json:"id"`
	Nickname string         `json:"nickname"`
	Version  *string        `json:"version"`
	Year     uint           `json:"year"`
	Brand    string         `json:"brand"`
	Model    string         `json:"model"`
	Photos   []VehiclePhoto `json:"photos"`
}

type VehiclePhoto struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}

type UserGroup struct {
	GroupID           uint    `json:"groupId"`
	GroupName         string  `json:"groupName"`
	GroupMainPhotoUrl *string `json:"groupMainPhotoUrl"`
}

type UserEvent struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	EventPhotoUrl *string `json:"eventPhotoUrl"`
}

type GetUsersResponse struct {
	Users       []GetUserResponseSimple `json:"users"`
	CurrentPage uint                    `json:"currentPage"`
	TotalPages  uint                    `json:"totalPages"`
}

type GetUserResponseSimple struct {
	ID            uint       `json:"id"`
	Name          string     `json:"name"`
	Nickname      string     `json:"nickname"`
	ProfilePicUrl *string    `json:"profilePicUrl"`
	Since         *time.Time `json:"since"`
}
