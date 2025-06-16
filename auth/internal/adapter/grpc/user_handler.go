package grpc

import (
	"context"
	useCase "go-grpc-clean/internal/domain/auth"
	pb "go-grpc-clean/internal/pb"
	"go-grpc-clean/internal/pkg/validator"
)

type AuthGRPCServer struct {
	pb.UnimplementedAuthServiceServer
	useCase useCase.IAuth
}

func NewGRPCServer(useCase useCase.IAuth) *AuthGRPCServer {
	return &AuthGRPCServer{useCase: useCase}
}

func (a *AuthGRPCServer) GetAnonymousToken(ctx context.Context, apiKey *pb.GetAnonymousTokenRequest) (*pb.GetAnonymousTokenResponse, error) {
	// validate
	params := useCase.GetAnonymousTokenRequest{
		ApiKey: apiKey.ApiKey,
	}

	if err := validator.Struct(params); err != nil {
		return nil, err
	}

	// business logic
	resp, err := a.useCase.GetAnonymousToken(ctx, params)
	if err != nil {
		return nil, err
	}

	// response
	return &pb.GetAnonymousTokenResponse{
		Token:           resp.Data.Token,
		TokenExpiration: resp.Data.TokenExpiration,
	}, nil
}
