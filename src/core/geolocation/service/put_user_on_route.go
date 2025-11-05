package service

import (
	"encoding/json"
	"fmt"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/infra/websocket/response"
	"github.com/redis/go-redis/v9"
)

func (gls *GeoLocationService) PutUserOnRoute(routeId, userId uint, data *geoinput.Coordinates) error {
	routeKey := "route:" + fmt.Sprint(routeId)

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
