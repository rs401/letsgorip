package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/forums/models"
	"github.com/rs401/letsgorip/forums/repository"
	"github.com/rs401/letsgorip/forums/service"
	"github.com/rs401/letsgorip/pb"
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

	forumsRepo := repository.NewForumsRepository(conn)
	// I think I'm going to seed forums here
	forums, err := forumsRepo.GetForums()
	if err != nil {
		log.Fatalf("Error checking if forums need to be seeded... : %v", err)
	}
	if len(forums) == 0 {
		// Seed the forums
		err := seedForums(forumsRepo)
		if err != nil {
			log.Fatalf("Error seeding forums: %v", err)
		}
	}
	forumService := service.NewForumService(forumsRepo)

	port, err := strconv.Atoi(os.Getenv("FORUMSVC_PORT"))
	if err != nil {
		log.Fatalf("Error getting forum service port: %v\n", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterForumServiceServer(grpcServer, forumService)
	log.Printf("Forum service running on port: :%d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error Serving: %v\n", err)
	}
}

var states []string = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

func seedForums(fr repository.ForumsRepository) error {
	seeds := createForums()
	for _, seed := range seeds {
		if err := fr.CreateForum(seed); err != nil {
			return err
		}
	}
	log.Println("Finished seeding Forums")
	// Seed Threads
	log.Println("Beginning seeding Threads")
	err := seedThreads(fr, seeds)
	if err != nil {
		return err
	}
	log.Println("Finished seeding Threads")
	return nil
}

func seedThreads(fr repository.ForumsRepository, seeds []*models.Forum) error {
	for _, forum := range seeds {
		thread := &models.Thread{
			ForumId: forum.Id,
			UserId:  forum.UserId,
			Title:   "Welcome to the forum.",
			Msg:     "This is the first thread in this forum, please take the time to say 'Hello 👋'.",
		}
		err := fr.CreateThread(thread)
		if err != nil {
			return err
		}
		// Seed a Post
		post := &models.Post{
			ThreadId: thread.Id,
			UserId:   forum.UserId,
			Msg:      "Hello 👋, 🤖.",
		}
		err = fr.CreatePost(post)
		if err != nil {
			return err
		}
	}
	return nil
}

func createForums() []*models.Forum {
	var seeds []*models.Forum = make([]*models.Forum, 0)
	for _, state := range states {
		var f *models.Forum = new(models.Forum)
		f.UserId = 0
		f.Title = state
		f.Description = fmt.Sprintf("Find people to off-road with in %s", state)
		seeds = append(seeds, f)
	}
	return seeds
}
