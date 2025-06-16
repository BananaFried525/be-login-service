package auth

import "context"

type IAuth interface {
	GetAnonymousToken(ctx context.Context, params GetAnonymousTokenRequest) (GetAnonymousTokenResponse, error)
}
