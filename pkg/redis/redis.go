package redis

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(redisURL string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opt)

	return rdb, nil
}
