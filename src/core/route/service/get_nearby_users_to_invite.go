package service

import (
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (s *RouteService) GetNearbyUsersToInvite(userId uint, data *input.GetNearbyUsersToInviteQuery) (*[]response.NearbyUserToRouteResponse, error) {
	if err := s.validator.Validate(data); err != nil {
		return nil, err
	}

	nearbyIDs, err := s.geoLocationService.GetNearbyUsersToRouteInvite(userId, data.Lat, data.Long, int(data.Page), int(data.Limit))
	if err != nil {
		return nil, err
	}

	if len(nearbyIDs) == 0 {
		return &[]response.NearbyUserToRouteResponse{}, nil
	}

	usersDetails, err := s.routeRepo.GetNearbyUsersDetails(nearbyIDs)
	if err != nil {
		return nil, generics.InternalError()
	}

	return usersDetails, nil
}
