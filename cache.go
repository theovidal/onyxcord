package onyxcord

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func OpenCache(config *Config) (client *redis.Client) {
	addr := config.Cache.Address + ":" + config.Cache.Port
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       0,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		log.Panicf("‼ Error connecting to cache: %s", err.Error())
	}

	return client
}
