package mocks

import "github.com/redis/go-redis/v9"

type AuthRedis struct {
	*redis.Client
}

func NewAuthRedis() *AuthRedis {
	return &AuthRedis{Client: redis.NewClient(&redis.Options{})}
}
