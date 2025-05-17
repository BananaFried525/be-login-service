package repository

import "context"

type UserRepository interface {
	Create(ctx context.Context, u *UserModel) (*UserModel, error)
	FindOne(ctx context.Context, u *UserModel) (*UserModel, error)
}
