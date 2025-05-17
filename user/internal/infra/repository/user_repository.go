package repository

import (
	"context"
	domain "go-grpc-clean/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func (useRepo *userRepositoryImpl) getInstance() *gorm.DB {
	instance := useRepo.db

	return instance
}

// Create implements repository.UserRepository.
func (useRepo *userRepositoryImpl) Create(ctx context.Context, u *domain.UserModel) (*domain.UserModel, error) {
	if err := useRepo.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (useRepo *userRepositoryImpl) FindOne(ctx context.Context, u *domain.UserModel) (*domain.UserModel, error) {
	instance := useRepo.getInstance()

	if u.ID != 0 {
		instance = instance.Where("id = ?", u.ID)
	}

	if u.Email != "" {
		instance = instance.Where("email = ?", u.Email)
	}

	if u.UserName != "" {
		instance = instance.Where("user_name = ?", u.UserName)
	}

	if u.ReferralCode != "" {
		instance = instance.Where("referral_code = ?", u.ReferralCode)
	}

	if err := instance.First(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func NewUserGormRepo(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{db: db}
}
