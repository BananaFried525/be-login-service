package usecase

import (
	"context"
	repository "go-grpc-clean/internal/domain/repository"
	domain "go-grpc-clean/internal/domain/user"
	"go-grpc-clean/internal/pb"
	"go-grpc-clean/internal/pkg/utils"
	"go-grpc-clean/internal/shared/response"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
)

type UserUseCaseImpl struct {
	repo repository.UserRepository
}

// CreateUser implements domain.IUserRepository.
func (u *UserUseCaseImpl) CreateUser(ctx context.Context, user *domain.User) (*domain.UserResponse, error) {
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

	createdAt := resp.CreatedAt.Format(time.RFC3339)
	updatedAt := resp.UpdatedAt.Format(time.RFC3339)

	status := map[repository.UserStatus]pb.UserStatus{
		repository.ACTIVE:    pb.UserStatus_ACTIVE,
		repository.INACTIVE:  pb.UserStatus_INACTIVE,
		repository.DELETED:   pb.UserStatus_DELETED,
		repository.SUSPENDED: pb.UserStatus_INACTIVE,
	}[resp.Status]

	return &domain.UserResponse{
		BaseResponse: response.BaseResponse[domain.User]{
			Status: response.OK,
			Data: domain.User{
				ID:           proto.Int32(resp.ID),
				ReferralCode: &resp.ReferralCode,
				UserName:     user.UserName,
				Email:        user.Email,
				Status:       &status,
				CreatedAt:    &createdAt,
				UpdatedAt:    &updatedAt,
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
