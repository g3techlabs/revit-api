package service

import (
	"encoding/json"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (s *RouteService) InviteNearbyUsers(userId, routeId uint, inviteds *input.UsersToInviteInput) error {
	if err := s.validator.Validate(inviteds); err != nil {
		return err
	}

	routeOwner, err := s.routeRepo.GetRouteOwner(userId, routeId)
	if err != nil {
		return generics.InternalError()
	}

	if !routeOwner.IsOwner {
		return generics.Forbidden("User is not owner of the route")
	}

	payload := response.RouteInvitePayload{
		RouteID: routeId,
		Inviter: routeOwner.Nickname,
		Destination: geoinput.Coordinates{
			Lat:  routeOwner.Lat,
			Long: routeOwner.Long,
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return generics.InternalError()
	}

	multicastMessage := websocket.MulticastMessage{
		TargetUserIDs: inviteds.IdsToInvite,
		Payload:       payloadBytes,
	}

	s.hub.Multicast <- &multicastMessage

	return nil
}
