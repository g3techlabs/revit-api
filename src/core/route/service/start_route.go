package service

import (
	"time"

	"github.com/g3techlabs/revit-api/src/core/route/errors"
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (s *RouteService) StartRoute(userId, routeId uint) error {
	if err := s.routeRepo.StartRoute(userId, routeId); err != nil {
		if err.Error() == "route not found" {
			return errors.RouteNotFound()
		}
		return generics.InternalError()
	}

	routeUsers, err := s.geoLocationService.GetUsersInRoute(routeId)
	if err != nil {
		return err
	}

	displayAt := time.Now().Add(5 * time.Second).Unix()
	payload := response.NewStartRouteEvent(routeId, displayAt)
	if err := s.hub.SendMulticastMessage(routeUsers, payload); err != nil {
		return err
	}

	return nil
}
