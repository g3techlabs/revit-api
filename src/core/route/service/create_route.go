package service

import (
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (rs *RouteService) CreateRoute(userId uint, data *input.CreateRouteInput) error {
	if err := rs.validator.Validate(data); err != nil {
		return err
	}

	routeId, err := rs.routeRepo.CreateRoute(userId, data.StartLocation, data.Destination)
	if err != nil {
		return generics.InternalError()
	}

	if err := rs.geoLocationService.PutUserOnRoute(routeId, userId, &data.StartLocation); err != nil {
		return err
	}

	return nil
}
