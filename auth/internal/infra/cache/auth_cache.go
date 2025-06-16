package cache

import (
	"context"
	"fmt"
	"go-grpc-clean/internal/domain/cache"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheImpl struct {
	redis *redis.Client
}

// GenerateKey implements cache.IAuth.
func (c CacheImpl) GenerateKey(params cache.GenerateKeyParams) string {
	return fmt.Sprintf("token:%s:%s:%s", params.Type, params.Role, params.Key)
}

// GetAnonymousToken implements cache.IAuth.
func (c CacheImpl) GetAnonymousToken(ctx context.Context, key string) (cache.TokenAttributes, error) {
	resp, err := c.redis.HMGet(ctx, key, "token").Result()
	if err != nil {
		return cache.TokenAttributes{}, err
	}

	_resp := cache.TokenAttributes{
		Key:    resp[0].(string),
		UserID: resp[1].(string),
		Role:   resp[2].(string),
		Scope:  resp[3].([]string),
		Expiry: resp[4].(string),
	}

	return _resp, err
}

func (c CacheImpl) SetAnonymousToken(ctx context.Context, key string, token interface{}, duration time.Duration) error {
	_, err := c.redis.Set(ctx, key, token, duration).Result()
	if err != nil {
		log.Printf("failed to set anonymous token: %v", err)
		return err
	}

	return nil
}

func NewAuth(redis *redis.Client) cache.IAuth {
	return CacheImpl{redis: redis}
}
