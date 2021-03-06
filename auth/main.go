// Package auth provides authentication/user API
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs401/letsgorip/auth/repository"
	"github.com/rs401/letsgorip/auth/service"
	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/pb"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading environment variables: %v", err)
	}
}

func main() {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	usersRepository := repository.NewUsersRepository(conn)
	authService := service.NewAuthService(usersRepository)
	// users, err := usersRepository.GetAll()
	// if err != nil {
	// 	log.Fatalf("Error retrieving users: %v\n", err)
	// }

	// fmt.Println(users)

	port, err := strconv.Atoi(os.Getenv("AUTHSVC_PORT"))
	if err != nil {
		log.Fatalf("Error getting auth service port: %v\n", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAuthServiceServer(grpcServer, authService)
	log.Printf("Auth service running on port: :%d\n", port)
	grpcServer.Serve(lis)
}
