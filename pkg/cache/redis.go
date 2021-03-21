package cache

import (
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"golang-gin-api/config"
)

// Create the Redis client
func InitRedisClient() (*redis.Client, error) {
	redisCfg := config.GetConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         redisCfg.Addr,
		Password:     redisCfg.Pass,
		DB:           redisCfg.Db,
		MaxRetries:   redisCfg.MaxRetries,
		PoolSize:     redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConns,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis error")
	}

	return client, nil
}
