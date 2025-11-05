package service

import (
	"encoding/json"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/infra/websocket/response"
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

	newPayload := &response.UserMovedEvent{
		Event: "user-moved",
		Payload: response.UserMovedPayload{
			Lat:    data.Lat,
			Lng:    data.Long,
			UserID: userId,
		},
	}

	payloadBytes, err := json.Marshal(newPayload)
	if err != nil {
		gls.logger.Errorf("Error while marshalling payload on PutUserOnFreeRoam: %v", err)
		return err
	}

	multicastMessage := &websocket.MulticastMessage{
		Payload:       payloadBytes,
		TargetUserIDs: targetIDs,
	}

	gls.hub.Multicast <- multicastMessage

	return nil
}
