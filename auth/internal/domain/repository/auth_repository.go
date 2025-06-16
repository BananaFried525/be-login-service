package repository

import "context"

type IAuth interface {
	GetToken(ctx context.Context, s *string) (*AuthModel, error)
}
