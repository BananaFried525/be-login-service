package cache

import (
	"context"
	"time"
)

type TokenAttributes struct {
	Key    string
	UserID string
	Role   string
	Scope  []string
	Expiry string
}

type GenerateKeyParams struct {
	Type string
	Role string
	Key  string
}

type IAuth interface {
	GenerateKey(params GenerateKeyParams) string
	SetAnonymousToken(ctx context.Context, key string, token interface{}, duration time.Duration) error
	GetAnonymousToken(ctx context.Context, key string) (TokenAttributes, error)
}
