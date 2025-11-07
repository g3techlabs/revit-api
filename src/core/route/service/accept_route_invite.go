package service

import (
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/route/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (s *RouteService) AcceptRouteInvite(userId, routeId uint, coordinates *geoinput.Coordinates) error {
	if err := s.validator.Validate(coordinates); err != nil {
		return err
	}

	userDetails, err := s.routeRepo.AcceptRouteInvite(userId, routeId, coordinates)
	if err != nil {
		switch err.Error() {
		case "route not found":
			return errors.RouteNotFound()
		case "user already participant":
			return errors.UserAlreadyInRoute("The user is already a participant of this route")
		}
		return generics.InternalError()
	}

	if err := s.geoLocationService.PutUserOnRoute(routeId, userDetails, coordinates); err != nil {
		return generics.InternalError()
	}

	return nil
}
