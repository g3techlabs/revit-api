package service

import (
	"fmt"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/geolocation/response"
	"github.com/redis/go-redis/v9"
)

func (gls *GeoLocationService) PutUserOnRoute(routeId uint, userDetails *response.UserDetails, data *geoinput.Coordinates) error {
	routeKey := "route:" + fmt.Sprint(routeId)

	userId := userDetails.UserId

	userCurrentKey, err := gls.geoLocationRepository.GetUserStateGeoKey(userId)
	if err != nil && err != redis.Nil {
		gls.logger.Errorf("Error getting user state geo key: %v", err)
		return err
	}

	if userCurrentKey == "" {
		gls.logger.Errorf("User %d (stateless) tried to join route %d", userId, routeId)
		return fmt.Errorf("forbidden action: user must be in free roam first")
	}

	if userCurrentKey != routeKey {
		if err := gls.geoLocationRepository.RemoveUserLocation(userCurrentKey, userId); err != nil {
			gls.logger.Errorf("Error removing user from key %s: %v", userCurrentKey, err)
			return err
		}
	}

	if err := gls.geoLocationRepository.SetUserState(routeKey, userId); err != nil {
		gls.logger.Errorf("Error setting user state: %v", err)
		return err
	}

	targetIDs, err := gls.geoLocationRepository.PutUserLocation(routeKey, userId, data)
	if err != nil {
		gls.logger.Errorf("Error while putting user on route: %v", err)
		return err
	}

	if len(targetIDs) == 0 {
		return nil
	}

	if err := gls.hub.SendMulticastMessage("new-user-in-route", targetIDs, *userDetails); err != nil {
		gls.logger.Errorf("Error marshalling UserAccepteedRouteInvite message: %v", err)
		return err
	}

	return nil
}
