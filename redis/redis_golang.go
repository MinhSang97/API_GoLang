package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisClient *redis.Client

func ConnectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Kiểm tra kết nối đến Redis
	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

}
