package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(host string, port int, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	redisClient := &RedisClient{}
	redisClient.Redis = rdb
	redisClient.Context = context.Background()

	return redisClient
}

type RedisClient struct {
	Redis   *redis.Client
	Context context.Context
}
