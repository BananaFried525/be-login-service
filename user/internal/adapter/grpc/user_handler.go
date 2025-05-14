package grpc

import (
	"context"
	domain "go-grpc-clean/internal/domain/user"
	pb "go-grpc-clean/internal/pb"
)

type GRPCServer struct {
	pb.UnimplementedUserServiceServer
	usecase domain.IUserRepository
}

func NewGRPCServer(u domain.IUserRepository) *GRPCServer {
	return &GRPCServer{usecase: u}
}

func (g *GRPCServer) CreateUser(ctx context.Context, user *pb.CreateUserRequest) (*pb.UserResponse, error) {
	resp, err := g.usecase.CreateUser(ctx, &domain.User{
		UserName: user.User.UserName,
		Email:    user.User.Email,
	})

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Status: pb.MessageStatus_OK,
		Data: &pb.User{
			Id:           resp.Data.ID,
			UserName:     resp.Data.UserName,
			Email:        resp.Data.Email,
			ReferralCode: resp.Data.ReferralCode,
			Status:       resp.Data.Status,
			CreatedAt:    resp.Data.CreatedAt,
			UpdatedAt:    resp.Data.UpdatedAt,
		},
	}, nil
}

func (g *GRPCServer) GetUser(ctx context.Context, user *pb.GetUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (g *GRPCServer) UpdateUser(ctx context.Context, user *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (g *GRPCServer) DeleteUser(ctx context.Context, user *pb.DeleteUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}
