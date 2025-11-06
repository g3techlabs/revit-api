package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func (s *GeoLocationRepository) GetNearbyUsersToRouteInvite(userId uint, lat, long float64, page, pageSize int) ([]uint, error) {
	ctx := context.Background()

	tempKey := fmt.Sprintf("temp:geo:invite:%d:%d", userId, time.Now().UnixNano())

	totalResults, err := s.redisClient.GeoSearchStore(ctx, FREE_ROAM_KEY, tempKey, &redis.GeoSearchStoreQuery{
		StoreDist: true,
		GeoSearchQuery: redis.GeoSearchQuery{
			Longitude:  long,
			Latitude:   lat,
			Radius:     float64(radiusInMetersOfRouteInvite),
			RadiusUnit: "m",
		},
	}).Result()

	if err != nil {
		return []uint{}, err
	}

	if totalResults == 0 {
		s.redisClient.Del(ctx, tempKey)
		return []uint{}, nil
	}

	start := int64((page - 1) * pageSize)
	stop := start + int64(pageSize) - 1

	resultStrings, err := s.redisClient.ZRange(ctx, tempKey, start, stop).Result()
	if err != nil {
		s.redisClient.Del(ctx, tempKey)
		return nil, err
	}

	s.redisClient.Del(ctx, tempKey)

	ids := s.convertNearbyUsersToIDs(resultStrings, userId)

	return ids, nil
}
