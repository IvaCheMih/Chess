package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client  *redis.Client
	context context.Context
}

func CreateRedisService(address string) RedisService {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	ctx := context.Background()
	return RedisService{
		client:  client,
		context: ctx,
	}
}
