package grpc

import (
	"context"
	domain "go-grpc-clean/internal/domain/user"
	pb "go-grpc-clean/internal/pb"
	"log"
	"time"

	protoLib "google.golang.org/protobuf/proto"
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
		log.Printf("failed to create user: %v", err)
		return nil, err
	}

	statusMap := map[domain.UserStatus]pb.UserStatus{
		domain.ACTIVE:    pb.UserStatus_ACTIVE,
		domain.INACTIVE:  pb.UserStatus_INACTIVE,
		domain.DELETED:   pb.UserStatus_DELETED,
		domain.SUSPENDED: pb.UserStatus_SUSPENDED,
	}

	status := statusMap[resp.Data.Status]
	createdAt := resp.Data.CreatedAt.Format(time.RFC3339)
	updatedAt := resp.Data.UpdatedAt.Format(time.RFC3339)

	return &pb.UserResponse{
		Status: pb.MessageStatus_OK,
		Data: &pb.User{
			Id:           protoLib.Int32(resp.Data.ID),
			UserName:     resp.Data.UserName,
			Email:        resp.Data.Email,
			ReferralCode: &resp.Data.ReferralCode,
			Status:       &status,
			CreatedAt:    &createdAt,
			UpdatedAt:    &updatedAt,
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
