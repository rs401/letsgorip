// Package api provides HTTP endpoints
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs401/letsgorip/api/handlers"
	"github.com/rs401/letsgorip/api/middlewares"
	"github.com/rs401/letsgorip/api/routes"
	"github.com/rs401/letsgorip/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authSvcHost  string
	forumSvcHost string
	authSvcPort  int
	forumSvcPort int
	apiPort      int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file (production?): %v\n", err)
	}
	authSvcHost = os.Getenv("AUTHSVC_HOST")
	authSvcPort, err = strconv.Atoi(os.Getenv("AUTHSVC_PORT"))
	if err != nil {
		log.Fatalf("Error converting AUTHSVC_PORT to int")
	}
	forumSvcHost = os.Getenv("FORUMSVC_HOST")
	forumSvcPort, err = strconv.Atoi(os.Getenv("FORUMSVC_PORT"))
	if err != nil {
		log.Fatalf("Error converting FORUMSVC_PORT to int")
	}
	apiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalf("Error converting API_PORT to int")
	}
}

func main() {
	// Client dial server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	authClientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", authSvcHost, authSvcPort), opts...)
	if err != nil {
		log.Fatal("Error dialing auth service:", err)
	}
	defer authClientConn.Close()
	forumClientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", forumSvcHost, forumSvcPort), opts...)
	if err != nil {
		log.Fatal("Error dialing forum service:", err)
	}
	defer forumClientConn.Close()
	// Initialize client
	// Pass this to handlers
	authSvcClient := pb.NewAuthServiceClient(authClientConn)
	forumSvcClient := pb.NewForumServiceClient(forumClientConn)
	// Setup handlers
	authHandleFuncs := handlers.NewAuthHandlers(authSvcClient)
	forumHandleFuncs := handlers.NewForumHandlers(forumSvcClient)
	// Setup router
	router := mux.NewRouter().StrictSlash(true)
	// Setup routes
	routes.SetupRoutes(router, authHandleFuncs, forumHandleFuncs)
	// Setup middlewares
	middlewares.SetupMiddleWares(router)
	// Listen
	log.Printf("Listening on port :%d\n", apiPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", apiPort), router)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
