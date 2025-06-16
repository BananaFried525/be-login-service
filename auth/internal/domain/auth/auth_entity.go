package auth

import (
	"go-grpc-clean/internal/shared/response"
)

type AnonymousAttributes struct {
	Token           string
	TokenExpiration string
}

type GetAnonymousTokenRequest struct {
	ApiKey string `json:"apiKey" validate:"required"`
}

type GetAnonymousTokenResponse struct {
	response.BaseResponse[AnonymousAttributes]
}
