package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var RedisClient *redis.Client
var Ctx=context.Background()

func ConnectRedis(){
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // Example: "localhost:6379"
		Password: "",                      // No password by default
		DB:       0,                        // Default DB
	})
	// Ping Redis to check connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}