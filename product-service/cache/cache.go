package cache

import (
	"context"
	"fmt"
	"product-service/configs"

	"github.com/go-redis/redis/v8"
)

func New(c configs.RedisConfiguration) *redis.Client {

	cache := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Username: c.Username,
		Password: c.Password, // no password set
		DB:       c.Db,       // use default DB
	})
	ctx := context.Background()
	pong, err := cache.Ping(ctx).Result()
	if err != nil {
		fmt.Println("connect redis fail: ", pong, err)
	}
	return cache
}
