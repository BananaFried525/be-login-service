package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpcAdapter "go-grpc-clean/internal/adapter/grpc"
	"go-grpc-clean/internal/infra/cache"
	"go-grpc-clean/internal/infra/redis"
	proto "go-grpc-clean/internal/pb"
	"go-grpc-clean/internal/pkg/config"
	useCase "go-grpc-clean/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.Init()
	configs := config.GetServiceConfig()

	// dbConn := database.NewGormDB()
	redisConn := redis.NewRedis()

	// userRepository := repository.New(dbConnection)
	cacheService := cache.NewAuth(redisConn)
	authUseCase := useCase.NewAuthUseCase(nil, cacheService)

	authGRPCServer := grpcAdapter.NewGRPCServer(authUseCase)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(recoveryUnaryInterceptor),
	)

	proto.RegisterAuthServiceServer(server, authGRPCServer)
	reflection.Register(server)

	log.Printf("server listening at %s\n", listener.Addr().String())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func recoveryUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[RECOVER] Unary panic recovered: %v", r)
			// err = grpc.Errorf(grpc.Code(grpc.Inter), "internal error")
		}
	}()
	return handler(ctx, req)
}
