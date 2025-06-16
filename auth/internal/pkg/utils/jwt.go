package utils

import (
	"go-grpc-clean/internal/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(data jwt.MapClaims) (tokenString string, err error) {
	secret := config.GetServiceConfig().JWTSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
