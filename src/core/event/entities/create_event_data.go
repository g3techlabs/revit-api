package entities

import "time"

type CreateEventData struct {
	Name         string
	Description  string
	Date         time.Time
	City         string
	VisibilityID uint
	Location     Coordinates
	GroupID      *uint
}

type Coordinates struct {
	Latitude  float64 `validate:"required,latitude"`
	Longitude float64 `validate:"required,longitude"`
}
