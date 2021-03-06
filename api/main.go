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
	"github.com/rs/cors"
	"github.com/rs401/letsgorip/api/handlers"
	"github.com/rs401/letsgorip/api/middlewares"
	"github.com/rs401/letsgorip/api/routes"
	"github.com/rs401/letsgorip/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	prod         string
	origins      []string
	authSvcHost  string
	forumSvcHost string
	placeSvcHost string
	authSvcPort  int
	forumSvcPort int
	placeSvcPort int
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
	placeSvcHost = os.Getenv("PLACESVC_HOST")
	placeSvcPort, err = strconv.Atoi(os.Getenv("PLACESVC_PORT"))
	if err != nil {
		log.Fatalf("Error converting PLACESVC_PORT to int")
	}
	apiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalf("Error converting API_PORT to int")
	}
	prod = os.Getenv("PROD")

}

func main() {
	if prod == "" {
		origins = []string{"http://localhost:4200"}
	} else {
		origins = []string{"http://letsgo.rip", "https://letsgo.rip"}
	}
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
	placeClientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", placeSvcHost, placeSvcPort), opts...)
	if err != nil {
		log.Fatal("Error dialing place service:", err)
	}
	defer placeClientConn.Close()
	// Initialize client
	// Pass this to handlers
	authSvcClient := pb.NewAuthServiceClient(authClientConn)
	forumSvcClient := pb.NewForumServiceClient(forumClientConn)
	placeSvcClient := pb.NewPlaceServiceClient(placeClientConn)
	// Setup handlers
	authHandleFuncs := handlers.NewAuthHandlers(authSvcClient)
	forumHandleFuncs := handlers.NewForumHandlers(forumSvcClient)
	placeHandleFuncs := handlers.NewPlaceHandlers(placeSvcClient)
	// Setup router
	router := mux.NewRouter().StrictSlash(true)
	// Setup routes
	routes.SetupRoutes(router, authHandleFuncs, forumHandleFuncs, placeHandleFuncs)
	// Setup middlewares
	middlewares.SetupMiddleWares(router)
	// Listen
	log.Printf("Listening on port :%d\n", apiPort)

	// corz := cors.AllowAll().Handler(router)
	corz := cors.New(cors.Options{
		AllowedOrigins: origins,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)
	// err = http.ListenAndServeTLS(fmt.Sprintf(":%d", apiPort), "./localhost.crt", "./localhost.key", corz)
	err = http.ListenAndServe(fmt.Sprintf(":%d", apiPort), corz)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
