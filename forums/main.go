package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/forums/repository"
	"github.com/rs401/letsgorip/forums/service"
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

	forumsRepo := repository.NewForumsRepository(conn)
	forumService := service.NewForumService(forumsRepo)
}
