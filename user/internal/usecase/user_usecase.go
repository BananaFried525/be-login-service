package usecase

import (
	"context"
	repository "go-grpc-clean/internal/domain/repository"
	domain "go-grpc-clean/internal/domain/user"
	"go-grpc-clean/internal/pkg/utils"
	"go-grpc-clean/internal/shared/response"
	"log"
)

type UserUseCaseImpl struct {
	repo repository.UserRepository
}

// CreateUser implements domain.IUserRepository.
func (u *UserUseCaseImpl) CreateUser(ctx context.Context, user *domain.User) (*domain.UserResponse, error) {
	log.Printf("creating user: %v", user)
	referralCode := utils.RandomString(12)

	payload := &repository.UserModel{
		ReferralCode: referralCode,
		UserName:     user.UserName,
		Email:        user.Email,
		Status:       repository.ACTIVE,
	}

	resp, err := u.repo.Create(ctx, payload)
	if err != nil {
		log.Printf("failed to create user: %v", err)

		return &domain.UserResponse{
			BaseResponse: response.BaseResponse[domain.User]{
				Status: response.BAD_REQUEST,
				Data:   domain.User{},
			},
		}, err
	}

	return &domain.UserResponse{
		BaseResponse: response.BaseResponse[domain.User]{
			Status: response.OK,
			Data: domain.User{
				ID:       resp.ID,
				UserName: user.UserName,
				Email:    user.Email,
			},
		},
	}, err
}

// DeleteUser implements domain.IUserRepository.
func (u *UserUseCaseImpl) DeleteUser(ctx context.Context, id string) (*domain.UserResponse, error) {
	panic("unimplemented")
}

// GetUser implements domain.IUserRepository.
func (u *UserUseCaseImpl) GetUser(ctx context.Context, id string) (*domain.UserResponse, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.IUserRepository.
func (u *UserUseCaseImpl) UpdateUser(ctx context.Context, user *domain.User) (*domain.UserResponse, error) {
	panic("unimplemented")
}

func NewUserUseCase(r repository.UserRepository) domain.IUserRepository {
	return &UserUseCaseImpl{repo: r}
}
