package repository

import "time"

type UserStatus string

const (
	ACTIVE    UserStatus = "ACTIVE"
	INACTIVE  UserStatus = "INACTIVE"
	DELETED   UserStatus = "DELETED"
	SUSPENDED UserStatus = "SUSPENDED"
)

type UserModel struct {
	ID           int32
	ReferralCode string
	UserName     string
	Email        string
	Status       UserStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (UserModel) TableName() string {
	return "users"
}
