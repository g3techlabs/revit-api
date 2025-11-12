package service

import (
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (s *RouteService) InviteUsers(userId, routeId uint, inviteds *input.UsersToInviteInput) error {
	if err := s.validator.Validate(inviteds); err != nil {
		return err
	}

	routeInfo, err := s.routeRepo.GetRouteAndOwnerInfo(userId, routeId)
	if err != nil {
		return generics.InternalError()
	}

	if !routeInfo.IsOwner {
		return generics.Forbidden("User is not owner of the route")
	}

	destination := geoinput.Coordinates{
		Lat:  routeInfo.Lat,
		Long: routeInfo.Long,
	}

	payload := response.NewRouteInvitePayload(routeId, destination, routeInfo.Nickname, routeInfo.ProfilePic)
	if err := s.hub.SendMulticastMessage("route-invite", inviteds.IdsToInvite, payload); err != nil {
		return generics.InternalError()
	}

	return nil
}
