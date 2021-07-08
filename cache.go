package onyxcord

import (
	"log"

	"github.com/go-redis/redis/v8"
)

// OpenCache returns the Redis cache client (if needed and taking credentials from the configuration file)
func OpenCache(config *Config) (client *redis.Client) {
	addr := config.Cache.Address + ":" + config.Cache.Port
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Cache.Password,
		DB:       0,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		log.Panicf("‼ Error connecting to cache: %s", err.Error())
	}

	return client
}
