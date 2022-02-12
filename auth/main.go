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

// const (
// 	serverCertFile   = "../cert/server-cert.pem"
// 	serverKeyFile    = "../cert/server-key.pem"
// 	clientCACertFile = "../cert/ca-cert.pem"
// )

// func loadTLSCredentials() (credentials.TransportCredentials, error) {
// 	// Load certificate of the CA who signed client's certificate
// 	pemClientCA, err := ioutil.ReadFile(clientCACertFile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	certPool := x509.NewCertPool()
// 	if !certPool.AppendCertsFromPEM(pemClientCA) {
// 		return nil, fmt.Errorf("failed to add client CA's certificate")
// 	}

// 	// Load server's certificate and private key
// 	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create the credentials and return it
// 	config := &tls.Config{
// 		Certificates: []tls.Certificate{serverCert},
// 		ClientAuth:   tls.RequireAndVerifyClientCert,
// 		ClientCAs:    certPool,
// 	}

// 	return credentials.NewTLS(config), nil
// }

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

	// tlsCredentials, err := loadTLSCredentials()
	// if err != nil {
	// 	log.Fatalf("cannot load TLS credentials: %v", err)
	// }

	var opts []grpc.ServerOption
	// opts = append(opts, grpc.Creds(tlsCredentials))

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAuthServiceServer(grpcServer, authService)
	log.Printf("Auth service running on port: :%d\n", port)
	grpcServer.Serve(lis)
}
