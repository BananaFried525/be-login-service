package domain

import "context"

type IUserRepository interface {
	CreateUser(ctx context.Context, user *User) (*UserResponse, error)
	GetUser(ctx context.Context, id string) (*UserResponse, error)
	UpdateUser(ctx context.Context, user *User) (*UserResponse, error)
	DeleteUser(ctx context.Context, id string) (*UserResponse, error)
}
