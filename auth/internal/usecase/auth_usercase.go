package usecase

import (
	"context"
	authDomain "go-grpc-clean/internal/domain/auth"
	"go-grpc-clean/internal/domain/cache"
	"go-grpc-clean/internal/domain/repository"
	"go-grpc-clean/internal/pkg/utils"
	"go-grpc-clean/internal/shared/response"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthUseCaseImpl struct {
	repository repository.IAuth
	cache      cache.IAuth
}

func NewAuthUseCase(repository repository.IAuth, cache cache.IAuth) authDomain.IAuth {
	return &AuthUseCaseImpl{
		repository: repository,
		cache:      cache,
	}
}

func (r *AuthUseCaseImpl) GetAnonymousToken(ctx context.Context, params authDomain.GetAnonymousTokenRequest) (authDomain.GetAnonymousTokenResponse, error) {
	key, err := uuid.NewV7()
	if err != nil {
		return authDomain.GetAnonymousTokenResponse{}, err
	}

	now := time.Now().Add(time.Hour * 12)
	exp := now.Format(time.RFC3339)

	tokenKey := r.cache.GenerateKey(cache.GenerateKeyParams{
		Type: "anonymous",
		Role: "anonymous",
		Key:  key.String(),
	})

	json, err := utils.JsonMarshal(cache.TokenAttributes{
		Key:    key.String(),
		UserID: key.String(),
		Role:   key.String(),
		Scope:  []string{key.String()},
		Expiry: exp,
	})
	if err != nil {
		return authDomain.GetAnonymousTokenResponse{}, err
	}

	err = r.cache.SetAnonymousToken(ctx, tokenKey, json, time.Duration(time.Hour*12))
	if err != nil {
		return authDomain.GetAnonymousTokenResponse{}, err
	}

	token, err := utils.GenerateJWT(jwt.MapClaims{
		"token":  tokenKey,
		"expiry": exp,
	})
	if err != nil {
		return authDomain.GetAnonymousTokenResponse{}, err
	}

	return authDomain.GetAnonymousTokenResponse{
		BaseResponse: response.BaseResponse[authDomain.AnonymousAttributes]{
			Status: response.OK,
			Data: authDomain.AnonymousAttributes{
				Token:           token,
				TokenExpiration: exp,
			},
		},
	}, nil
}
