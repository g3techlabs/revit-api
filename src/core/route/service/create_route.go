package service

import (
	georesponse "github.com/g3techlabs/revit-api/src/core/geolocation/response"
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (rs *RouteService) CreateRoute(userId uint, data *input.CreateRouteInput) (*response.RouteCreatedReponse, error) {
	if err := rs.validator.Validate(data); err != nil {
		return nil, err
	}

	routeId, err := rs.routeRepo.CreateRoute(userId, data.StartLocation, data.Destination)
	if err != nil {
		return nil, generics.InternalError()
	}

	userDetails := georesponse.UserDetails{
		UserId: userId,
	}

	if err := rs.geoLocationService.PutUserOnRoute(routeId, &userDetails, &data.StartLocation); err != nil {
		return nil, err
	}

	return &response.RouteCreatedReponse{RouteID: routeId}, nil
}
