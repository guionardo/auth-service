package data

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	CacheRedis struct {
		client *redis.Client
	}
)

var ctx = context.Background()

func (c *CacheRedis) Setup(args interface{}) Cache {
	if options, ok := args.(redis.Options); ok {
		c.client = redis.NewClient(&options)
	} else if client, ok := args.(*redis.Client); ok {
		c.client = client
	} else {
		panic("Expected redis.Options or redis.Client for CacheRedis.Setup argument")
	}

	return c
}

func (c *CacheRedis) Set(key string, value interface{}, timeToLive time.Duration) error {
	status := c.client.Set(ctx, key, value, timeToLive)
	return status.Err()
}

func (c *CacheRedis) Get(key string) (interface{}, error) {
	val, err := c.client.Get(ctx, key).Result()
	return val, err
}
