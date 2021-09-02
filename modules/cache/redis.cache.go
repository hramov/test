package cache

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Cache struct {
	Client *redis.Client
}

func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	log.Println("Redis has been started")
	return rdb
}

func (c *Cache) Set(key string, value string) (string, error) {
	err := c.Client.Set(ctx, key, value, 1).Err()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return value, nil
}

func (c *Cache) Get(key string) (string, error) {
	value, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return value, nil
}
