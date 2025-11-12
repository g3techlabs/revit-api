package service

import (
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/geolocation/response"
	"github.com/redis/go-redis/v9"
)

func (gls *GeoLocationService) PutUserOnFreeRoam(userId uint, data *geoinput.Coordinates) error {
	if err := gls.validator.Validate(data); err != nil {
		gls.logger.Errorf("%v", err)
		return err
	}

	userCurrentKey, err := gls.geoLocationRepository.GetUserStateGeoKey(userId)
	if err != nil && err != redis.Nil {
		gls.logger.Errorf("Error getting user state geo key: %v", err)
		return err
	}

	if userCurrentKey != "" && userCurrentKey != FREE_ROAM_KEY {
		if err := gls.geoLocationRepository.RemoveUserLocation(userCurrentKey, userId); err != nil {
			gls.logger.Errorf("Error setting new user state: %v", err)
			return err
		}
	}

	if err := gls.geoLocationRepository.SetUserState(FREE_ROAM_KEY, userId); err != nil {
		gls.logger.Errorf("Error setting user state to free roam: %v", err)
		return err
	}

	targetIDs, err := gls.geoLocationRepository.PutUserLocation(FREE_ROAM_KEY, userId, data)
	if err != nil {
		gls.logger.Errorf("Error while putting user on free roam: %v", err)
		return err
	}

	if len(targetIDs) == 0 {
		return nil
	}

	payload := response.NewUserMovedPayload(userId, data)
	if err := gls.hub.SendMulticastMessage("user-moved", targetIDs, payload); err != nil {
		gls.logger.Errorf("Error marshalling UserMoved message: %v", err)
		return err
	}

	return nil
}
