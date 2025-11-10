package response

import geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"

type RouteInviteEvent struct {
	Event   string             `json:"event"`
	Payload RouteInvitePayload `json:"payload"`
}
type RouteInvitePayload struct {
	RouteID        uint                 `json:"routeId"`
	Destination    geoinput.Coordinates `json:"destination"`
	InviterDetails InviterDetails       `json:"inviter"`
}

type InviterDetails struct {
	InviterName       string  `json:"inviterName"`
	InviterProfilePic *string `json:"inviterProfilePic"`
}

func NewRouteInviteEvent(routeId uint, coordinates geoinput.Coordinates, inviterName string, inviterProfilePic *string) RouteInviteEvent {
	inviterDetails := InviterDetails{
		InviterName:       inviterName,
		InviterProfilePic: inviterProfilePic,
	}

	return RouteInviteEvent{
		Event: "route-invite",
		Payload: RouteInvitePayload{
			RouteID:        routeId,
			Destination:    coordinates,
			InviterDetails: inviterDetails,
		},
	}
}
