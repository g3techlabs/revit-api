package response

import (
	"time"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
)

type UserMovedEvent struct {
	Event   string           `json:"event"`
	Payload UserMovedPayload `json:"payload"`
}

type UserMovedPayload struct {
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	UserID    uint      `json:"userId"`
	Timestamp time.Time `json:"timestamp"`
}

func NewUserMovedEvent(userId uint, coordinates *geoinput.Coordinates) *UserMovedEvent {
	return &UserMovedEvent{
		Event: "user-moved",
		Payload: UserMovedPayload{
			Lat:       coordinates.Lat,
			Lng:       coordinates.Long,
			UserID:    userId,
			Timestamp: time.Now().UTC(),
		},
	}
}
