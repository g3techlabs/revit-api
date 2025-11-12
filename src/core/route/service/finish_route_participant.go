package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/route/response"
)

func (s *RouteService) FinishRouteParticipant(userId uint, coordinates *geoinput.Coordinates) error {
	if err := s.validator.Validate(coordinates); err != nil {
		return err
	}

	routeKey, err := s.geoLocationService.GetUserCurrentKey(userId)
	if err != nil {
		return err
	}

	routeId, err := s.parseRouteIdFromKey(routeKey)
	if err != nil {
		return err
	}

	finishDetails, err := s.routeRepo.FinishParticipant(userId, routeId, time.Now())
	if err != nil {
		return err
	}

	payload := response.NewUserFinishedRoutePayload(userId, finishDetails.ArrivalTime, finishDetails.ArrivalOrder)
	if err := s.hub.SendSinglecastMessage("user-finished-route", userId, payload); err != nil {
		return err
	}

	if err := s.geoLocationService.PutUserOnFreeRoam(userId, coordinates); err != nil {
		return err
	}

	return nil
}

func (s *RouteService) parseRouteIdFromKey(routeKey string) (uint, error) {
	stringRouteId, ok := strings.CutPrefix(routeKey, "route:")
	if !ok {
		return 0, fmt.Errorf("invalid route key")
	}

	routeId, err := strconv.ParseUint(stringRouteId, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse route id")
	}

	return uint(routeId), nil
}
