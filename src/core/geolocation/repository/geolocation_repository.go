package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/g3techlabs/revit-api/src/config"
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/geolocation/response"
	"github.com/redis/go-redis/v9"
)

const (
	USER_STATE_KEY_PREFIX        = "user:state:"
	GEO_KEY_FIELD                = "current_geo_key"
	FREE_ROAM_KEY         string = "free-roam"
)

var nearbyUsersRadiusInKm = config.GetIntVariable("USERS_RADIUS_IN_KM")
var radiusInMetersOfRouteInvite = config.GetIntVariable("RADIUS_IN_METERS_OF_ROUTE_INVITE")

type IGeoLocationRepository interface {
	PutUserLocation(key string, userId uint, data *geoinput.Coordinates) ([]uint, error)
	GetFreeRoamNearbyUsers(key string, userId uint, coordinates *geoinput.Coordinates) (*[]response.NearbyUsers, error)
	RemoveUserLocation(key string, userId uint) error
	GetUserStateGeoKey(userId uint) (string, error)
	SetUserState(key string, userId uint) error
	ClearUserState(userId uint) error
}

type GeoLocationRepository struct {
	redisClient *redis.Client
}

func NewGeoLocationRepository(redisClient *redis.Client) IGeoLocationRepository {
	return &GeoLocationRepository{
		redisClient: redisClient,
	}
}

func (s *GeoLocationRepository) userStateKey(userId uint) string {
	return fmt.Sprintf("%s%d", USER_STATE_KEY_PREFIX, userId)
}

func (s *GeoLocationRepository) GetUserStateGeoKey(userId uint) (string, error) {
	ctx := context.Background()

	return s.redisClient.HGet(ctx, s.userStateKey(userId), GEO_KEY_FIELD).Result()
}

func (s *GeoLocationRepository) SetUserState(key string, userId uint) error {
	ctx := context.Background()

	return s.redisClient.HSet(ctx, s.userStateKey(userId), GEO_KEY_FIELD, key).Err()
}

func (s *GeoLocationRepository) ClearUserState(userId uint) error {
	ctx := context.Background()

	return s.redisClient.Del(ctx, s.userStateKey(userId)).Err()
}

func (s *GeoLocationRepository) GetFreeRoamNearbyUsers(key string, userId uint, data *geoinput.Coordinates) (*[]response.NearbyUsers, error) {
	ctx := context.Background()
	res, err := s.redisClient.GeoSearchLocation(ctx, key, &redis.GeoSearchLocationQuery{
		WithCoord: true,
		GeoSearchQuery: redis.GeoSearchQuery{
			Longitude:  data.Long,
			Latitude:   data.Lat,
			Radius:     float64(radiusInMetersOfRouteInvite),
			RadiusUnit: "m",
		},
	}).Result()

	if err != nil {
		return nil, err
	}

	nearbyUsers := s.convertToNearbyUsers(res, userId)

	return nearbyUsers, nil
}

func (s *GeoLocationRepository) convertToNearbyUsers(res []redis.GeoLocation, requesterId uint) *[]response.NearbyUsers {
	var nearbyUsers []response.NearbyUsers
	for _, user := range res {
		userId, err := strconv.ParseUint(user.Name, 10, 64)
		if err != nil {
			continue
		}

		if userId == uint64(requesterId) {
			continue
		}

		nearbyUsers = append(nearbyUsers, response.NearbyUsers{
			UserID: uint(userId),
			Lat:    user.Latitude,
			Lng:    user.Longitude,
		})
	}

	return &nearbyUsers
}

func (s *GeoLocationRepository) convertNearbyUsersToIDs(res []string, senderId uint) []uint {
	var targetIDs []uint
	for _, userIdStr := range res {
		userId, err := strconv.ParseUint(userIdStr, 10, 64)
		if err != nil {
			continue
		}

		if uint(userId) == senderId {
			continue
		}

		targetIDs = append(targetIDs, uint(userId))
	}

	return targetIDs
}

func (s *GeoLocationRepository) RemoveUserLocation(key string, userId uint) error {
	ctx := context.Background()

	_, err := s.redisClient.ZRem(ctx, key, fmt.Sprint(userId)).Result()
	if err != nil {
		return err
	}

	return nil
}
