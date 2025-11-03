package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/core/geolocation/input"
	"github.com/g3techlabs/revit-api/src/core/geolocation/response"
	"github.com/redis/go-redis/v9"
)

type IGeoLocationRepository interface {
	PutUserLocation(userId uint, data *input.Coordinates) ([]uint, error)
	GetNearbyUsers(userId uint, coordinates *input.Coordinates) (*[]response.NearbyUsers, error)
	RemoveUserLocation(userId uint) error
}

type GeoLocationRepository struct {
	ctx         context.Context
	redisClient *redis.Client
}

func NewGeoLocationRepository(ctx context.Context, redisClient *redis.Client) IGeoLocationRepository {
	return &GeoLocationRepository{
		ctx:         ctx,
		redisClient: redisClient,
	}
}

var nearbyUsersRadiusInKm = config.GetIntVariable("USERS_RADIUS_IN_KM")

func (s *GeoLocationRepository) PutUserLocation(userId uint, data *input.Coordinates) ([]uint, error) {
	_, err := s.redisClient.GeoAdd(s.ctx, "users:location", &redis.GeoLocation{
		Longitude: data.Long,
		Latitude:  data.Lat,
		Name:      fmt.Sprint(userId),
	}).Result()

	if err != nil {
		return nil, err
	}

	targetIDs, err := s.getNearbyUsersIds(data.Lat, data.Long)
	if err != nil {
		return nil, err
	}

	return targetIDs, nil
}

func (s *GeoLocationRepository) getNearbyUsersIds(lat, lng float64) ([]uint, error) {
	res, err := s.redisClient.GeoSearchLocation(s.ctx, "users:location", &redis.GeoSearchLocationQuery{
		WithCoord: false,
		GeoSearchQuery: redis.GeoSearchQuery{
			Longitude:  lng,
			Latitude:   lat,
			Radius:     float64(nearbyUsersRadiusInKm),
			RadiusUnit: "km",
		},
	}).Result()

	if err != nil {
		return nil, err
	}

	targetIds := s.convertNearbyUsersToIDs(res)
	return targetIds, nil
}

func (s *GeoLocationRepository) GetNearbyUsers(userId uint, data *input.Coordinates) (*[]response.NearbyUsers, error) {
	res, err := s.redisClient.GeoSearchLocation(s.ctx, "users:location", &redis.GeoSearchLocationQuery{
		WithCoord: true,
		GeoSearchQuery: redis.GeoSearchQuery{
			Longitude:  data.Long,
			Latitude:   data.Lat,
			Radius:     float64(nearbyUsersRadiusInKm),
			RadiusUnit: "km",
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

func (s *GeoLocationRepository) convertNearbyUsersToIDs(res []redis.GeoLocation) []uint {
	var targetIDs []uint
	for _, userNearby := range res {
		userId, err := strconv.ParseUint(userNearby.Name, 10, 64)
		if err != nil {
			continue
		}
		targetIDs = append(targetIDs, uint(userId))
	}

	return targetIDs
}

func (s *GeoLocationRepository) RemoveUserLocation(userId uint) error {
	_, err := s.redisClient.ZRem(s.ctx, "users:location", fmt.Sprint(userId)).Result()
	if err != nil {
		return err
	}

	return nil
}
