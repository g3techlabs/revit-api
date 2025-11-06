package response

import geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"

type RouteInvitePayload struct {
	RouteID     uint                 `json:"routeId"`
	Destination geoinput.Coordinates `json:"destination"`
	Inviter     string               `json:"inviter"`
}
