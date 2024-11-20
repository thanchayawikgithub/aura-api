package cache

import (
	"aura/internal/config"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	IRedisClient interface {
		Set(ctx context.Context, key string, value interface{}) error
		Get(ctx context.Context, key string, value interface{}) error
		Del(ctx context.Context, keys ...string) error
		IsExist(ctx context.Context, key string) (bool, error)
		Keys(ctx context.Context, key string) ([]string, error)
	}

	RedisClient struct {
		Client   *redis.Client
		Duration time.Duration
	}
)

func NewRedisClient(cfg config.Redis) *RedisClient {
	return &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Password: "",
			DB:       cfg.DB,
		}),
		Duration: time.Minute * 15,
	}
}

func (c *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Client.Set(ctx, key, bytes, c.Duration).Err()
}

func (c *RedisClient) Get(ctx context.Context, key string, value interface{}) error {
	bytes, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, value)
}

func (c *RedisClient) Del(ctx context.Context, keys ...string) error {
	return c.Client.Del(ctx, keys...).Err()
}

func (c *RedisClient) IsExist(ctx context.Context, key string) (bool, error) {
	result, err := c.Client.Exists(ctx, key).Result()
	return result > 0, err
}

func (c *RedisClient) Keys(ctx context.Context, key string) ([]string, error) {
	return c.Client.Keys(ctx, key).Result()
}
