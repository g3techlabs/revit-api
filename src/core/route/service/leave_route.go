package service

import (
	"time"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/route/response"
)

func (s *RouteService) LeaveRoute(userId uint, coordinates *geoinput.Coordinates) error {
	routeKey, err := s.geoLocationService.GetUserCurrentKey(userId)
	if err != nil {
		return err
	}

	routeId, err := s.parseRouteIdFromKey(routeKey)
	if err != nil {
		return err
	}

	if err := s.routeRepo.LeaveRoute(userId, routeId, time.Now()); err != nil {
		return err
	}

	if err := s.geoLocationService.PutUserOnFreeRoam(userId, coordinates); err != nil {
		return err
	}

	remainingParticipantIDs, err := s.geoLocationService.GetUsersInRoute(routeId)
	if err != nil {
		return err
	}

	if len(remainingParticipantIDs) > 0 {
		payload := response.NewParticipantLeftRoutePayload(userId)
		if err := s.hub.SendMulticastMessage("participant-left-route", remainingParticipantIDs, payload); err != nil {
			return err
		}
	}

	return nil
}
