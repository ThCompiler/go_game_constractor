package redis

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(redisUrl string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opt)
	return rdb, nil
}
