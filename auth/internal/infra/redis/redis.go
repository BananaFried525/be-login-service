package redis

import (
	"context"
	"fmt"
	"go-grpc-clean/internal/pkg/config"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	configs := config.GetRedisConfig()

	db, err := strconv.Atoi(configs.DB)
	if err != nil {
		log.Fatalf("failed to convert db to int: %v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.Host, configs.Port),
		DB:       db,
		Password: configs.Password,
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to ping redis: %v\n", err)
	}

	log.Println("connected to redis")

	return client
}
