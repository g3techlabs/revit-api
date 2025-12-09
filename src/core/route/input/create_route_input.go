package input

import geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"

type CreateRouteInput struct {
	StartLocation geoinput.Coordinates `json:"startLocation" validate:"required"`
	Destination   geoinput.Coordinates `json:"destination" validate:"required"`
}
