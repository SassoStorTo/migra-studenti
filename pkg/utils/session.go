package utils

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var redisClient *redis.Client

func InitStoreSess() {
	// Todo: set passwd and things
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Redis server address
		Password: "",           // No password
		DB:       0,            // Use default DB
	})
}

func SetStore(key string, value any, exp_time time.Duration, c *fiber.Ctx) error {
	return redisClient.Set(c.Context(), key, value, exp_time).Err()
}

func GetValue(key string, result interface{}, c *fiber.Ctx) error {
	res, err := redisClient.Get(c.Context(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(res), result)
}

func ResetValue(key string, c *fiber.Ctx) (int64, error) {
	return redisClient.Del(c.Context(), key).Result()
}

func StoreRoute(c *fiber.Ctx) error {
	return SetStore("original-route", c.Route().Path, time.Minute*5, c)
}
