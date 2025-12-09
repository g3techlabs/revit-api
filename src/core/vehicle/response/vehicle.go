package response

import "time"

type SimpleVehicle struct {
	ID           uint    `json:"id"`
	Nickname     string  `json:"nickname"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Year         uint    `json:"year"`
	Version      *string `json:"version"`
	MainPhotoUrl *string `json:"mainPhotoUrl"`
}

type Photo struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}

type VehiclePhoto struct {
	ID        uint   `json:"id"`
	Reference string `json:"reference"`
	VehicleID uint   `json:"vehicleId"`
}

type GetVehicleResponse struct {
	Nickname  string         `json:"nickname"`
	Brand     string         `json:"brand"`
	Model     string         `json:"model"`
	Year      uint           `json:"year"`
	Version   *string        `json:"version"`
	MainPhoto *string        `json:"mainPhoto"`
	CreatedAt time.Time      `json:"createdAt"`
	Photos    []VehiclePhoto `json:"photos"`
}

type GetVehiclesResponse struct {
	Vehicles    []SimpleVehicle `json:"vehicles"`
	CurrentPage uint            `json:"currentPage"`
	TotalPages  uint            `json:"totalPages"`
}
