package response

import geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"

type RouteInvitePayload struct {
	RouteID        uint                 `json:"routeId"`
	Destination    geoinput.Coordinates `json:"destination"`
	InviterDetails InviterDetails       `json:"inviter"`
}

type InviterDetails struct {
	InviterName       string  `json:"inviterName"`
	InviterProfilePic *string `json:"inviterProfilePic"`
}

func NewRouteInvitePayload(routeId uint, coordinates geoinput.Coordinates, inviterName string, inviterProfilePic *string) RouteInvitePayload {
	inviterDetails := InviterDetails{
		InviterName:       inviterName,
		InviterProfilePic: inviterProfilePic,
	}

	return RouteInvitePayload{
		RouteID:        routeId,
		Destination:    coordinates,
		InviterDetails: inviterDetails,
	}
}
