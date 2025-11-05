package repository

import (
	"context"
	"fmt"

	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/redis/go-redis/v9"
)

func (s *GeoLocationRepository) PutUserLocation(key string, userId uint, data *geoinput.Coordinates) ([]uint, error) {
	ctx := context.Background()

	_, err := s.redisClient.GeoAdd(ctx, key, &redis.GeoLocation{
		Longitude: data.Long,
		Latitude:  data.Lat,
		Name:      fmt.Sprint(userId),
	}).Result()

	if err != nil {
		return nil, err
	}

	targetIDs, err := s.getTargetIds(key, userId, data.Lat, data.Long)
	if err != nil {
		return nil, err
	}

	return targetIDs, nil
}

func (s *GeoLocationRepository) getTargetIds(key string, senderId uint, lat, lng float64) ([]uint, error) {
	var res []string
	var err error

	if key != FREE_ROAM_KEY {
		res, err = s.getUsersWithinRoute(key)
		if err != nil {
			return nil, err
		}
	} else {
		res, err = s.getFreeRoamUsersIds(key, lat, lng)
		if err != nil {
			return nil, err
		}
	}

	targetIds := s.convertNearbyUsersToIDs(res, senderId)
	return targetIds, nil
}

func (s *GeoLocationRepository) getUsersWithinRoute(key string) ([]string, error) {
	ctx := context.Background()

	return s.redisClient.ZRange(ctx, key, 0, -1).Result()
}

func (s *GeoLocationRepository) getFreeRoamUsersIds(key string, lat, lng float64) ([]string, error) {
	ctx := context.Background()

	return s.redisClient.GeoSearch(ctx, key, &redis.GeoSearchQuery{
		Longitude:  lng,
		Latitude:   lat,
		Radius:     float64(nearbyUsersRadiusInKm),
		RadiusUnit: "km",
	}).Result()
}
