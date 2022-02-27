// Package places provides the Places API
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/pb"
	"github.com/rs401/letsgorip/places/repository"
	"github.com/rs401/letsgorip/places/service"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading environment variables, (production?): %v", err)
	}
}

func main() {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	placesRepo := repository.NewPlacesRepository(conn)
	placesSvc := service.NewPlaceService(placesRepo)

	port, err := strconv.Atoi(os.Getenv("PLACESVC_PORT"))
	if err != nil {
		log.Fatalf("Error getting place service port: %v\n", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPlaceServiceServer(grpcServer, placesSvc)
	log.Printf("Forum service running on port: :%d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error Serving: %v\n", err)
	}

}
