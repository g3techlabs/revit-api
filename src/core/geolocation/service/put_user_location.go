package service

import (
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/geolocation/response"
	"github.com/redis/go-redis/v9"
)

func (gls *GeoLocationService) PutUserLocation(userId uint, data *geoinput.Coordinates) error {
	if err := gls.validator.Validate(data); err != nil {
		gls.logger.Errorf("%v", err)
		return err
	}

	key, err := gls.geoLocationRepository.GetUserStateGeoKey(userId)
	if err != nil && err != redis.Nil {
		gls.logger.Errorf("Error getting user state geo key: %v", err)
		return err
	}

	isNewState := false
	if key == "" {
		key = FREE_ROAM_KEY
		isNewState = true
	}

	if isNewState {
		if err := gls.geoLocationRepository.SetUserState(key, userId); err != nil {
			gls.logger.Errorf("Error setting new user state to free roam: %v", err)
			return err
		}
	}

	targetIDs, err := gls.geoLocationRepository.PutUserLocation(key, userId, data)
	if err != nil {
		gls.logger.Errorf("Error while putting user location: %v", err)
		return err
	}

	payload := response.NewUserMovedEvent(userId, data)
	if err := gls.hub.SendMulticastMessage(targetIDs, payload); err != nil {
		gls.logger.Errorf("Error marshalling UserMovedEvent message: %v", err)
		return err
	}
	return nil
}
