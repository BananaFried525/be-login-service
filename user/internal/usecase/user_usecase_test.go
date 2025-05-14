package usecase_test

import (
	"context"
	"testing"

	domain "go-grpc-clean/internal/domain/user"
	"go-grpc-clean/internal/usecase"
	"go-grpc-clean/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func setupUserUseCase() (domain.IUserRepository, *mocks.UserRepository) {
	repo := new(mocks.UserRepository)
	userUseCase := usecase.NewUserUseCase(repo)
	return userUseCase, repo
}

func TestGetUser_Success(t *testing.T) {
	userUseCase, _ := setupUserUseCase()
	ctx := context.Background()

	user, err := userUseCase.GetUser(ctx, "123")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "123", user.Data.ID)
}

func TestGetUser_Failure_EmptyID(t *testing.T) {
	userUseCase, _ := setupUserUseCase()
	ctx := context.Background()

	user, err := userUseCase.GetUser(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "id_is_empty", err.Error())
}
