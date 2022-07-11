package infrastructure

import (
	"fmt"

	"github.com/go-redis/redis/v9"
)

func NewRedis(config *Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	return rdb
}
