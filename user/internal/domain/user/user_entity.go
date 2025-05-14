package domain

import (
	"go-grpc-clean/internal/pb"
	"go-grpc-clean/internal/shared/response"
)

type User struct {
	ID           *int32
	ReferralCode *string
	UserName     string
	Email        string
	Status       *pb.UserStatus
	CreatedAt    *string
	UpdatedAt    *string
}

type UserResponse struct {
	response.BaseResponse[User]
}
