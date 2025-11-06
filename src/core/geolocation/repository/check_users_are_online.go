package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (s *GeoLocationRepository) CheckUsersAreOnline(userIDs []uint) ([]bool, error) {
	ctx := context.Background()

	pipe := s.redisClient.Pipeline()

	cmds := make([]*redis.IntCmd, len(userIDs))

	for i, id := range userIDs {
		key := s.userStateKey(id)
		cmds[i] = pipe.Exists(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]bool, len(userIDs))
	for i, cmd := range cmds {
		results[i] = cmd.Val() == 1
	}

	return results, nil
}
