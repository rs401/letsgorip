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
	"github.com/rs401/letsgorip/authapi/handlers"
	"github.com/rs401/letsgorip/authapi/middlewares"
	"github.com/rs401/letsgorip/authapi/routes"
	"github.com/rs401/letsgorip/pb"
	"google.golang.org/grpc"
)

var (
	authSvcHost string
	authSvcPort int
	apiPort     int
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
	apiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalf("Error converting API_PORT to int")
	}
}
func main() {
	// Client dial server
	clientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", authSvcHost, authSvcPort))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer clientConn.Close()
	// Initialize client
	// Pass this to handlers
	// authSvcClient := authclient.NewAuthServiceClient(client)
	authSvcClient := pb.NewAuthServiceClient(clientConn)
	// Setup handlers
	hndFuncs := handlers.NewAuthHandlers(authSvcClient)
	// Setup router
	router := mux.NewRouter()
	// Setup routes
	routes.SetupRoutes(router, hndFuncs)
	// Setup middlewares
	middlewares.SetupMiddleWares(router)
	// Listen
	log.Printf("Listening on port :%d\n", apiPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", apiPort), router)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
