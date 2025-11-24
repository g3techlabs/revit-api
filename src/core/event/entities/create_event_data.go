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

// Coordinates representa coordenadas geográficas para eventos
// @Description Coordenadas geográficas (latitude e longitude) para localização de eventos
type Coordinates struct {
	// Latitude da localização do evento
	Latitude float64 `validate:"required,latitude" example:"-23.5505"`
	// Longitude da localização do evento
	Longitude float64 `validate:"required,longitude" example:"-46.6333"`
}
