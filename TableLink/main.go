package main

import (
	"log"
	"net"

	db "tablelink/db/sql"
	"tablelink/middleware"
	"tablelink/proto/auth"
	"tablelink/proto/users"
	"tablelink/repository"
	"tablelink/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// Buat gRPC server dengan middleware AuthMiddleware
	grpcServerMw := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthMiddleware("be")),
	)
	db :=db.DB
	authRepo := repository.NewAuthRepository(db)
	userRepo := repository.NewUserRepository(db)

	auth.RegisterAuthServiceServer(grpcServer, server.NewAuthServer(authRepo))
	users.RegisterUserServiceServer(grpcServerMw, server.NewUserServer(userRepo))

	// // Register Auth Service TANPA Middleware
	// authRepo := repository.NewAuthRepository()
	// auth.RegisterAuthServiceServer(grpcServer, server.NewAuthHandler(authRepo))

	// // Register User Service DENGAN Middleware
	// userRepo := repository.NewUserRepository()
	// userServer := server.NewUserServer(userRepo)

	// userService := grpc.NewServer(
	// 	grpc.UnaryInterceptor(middleware.AuthMiddleware("be")), // Middleware khusus untuk user
	// )
	// users.RegisterUserServiceServer(userService, userServer)

	log.Println("gRPC server listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
