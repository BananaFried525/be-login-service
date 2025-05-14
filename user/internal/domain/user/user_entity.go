package domain

import (
	"time"

	"go-grpc-clean/internal/shared/response"
)

type UserStatus string

const (
	ACTIVE    UserStatus = "ACTIVE"
	INACTIVE  UserStatus = "INACTIVE"
	DELETED   UserStatus = "DELETED"
	SUSPENDED UserStatus = "SUSPENDED"
)

type User struct {
	ID           int32
	ReferralCode string
	UserName     string
	Email        string
	Status       UserStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserResponse struct {
	response.BaseResponse[User]
}
