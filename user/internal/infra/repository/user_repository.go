package repository

import (
	"context"
	domain "go-grpc-clean/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

// Create implements repository.UserRepository.
func (userRep *userRepositoryImpl) Create(ctx context.Context, u *domain.UserModel) (*domain.UserModel, error) {
	if err := userRep.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func NewUserGormRepo(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{db: db}
}
