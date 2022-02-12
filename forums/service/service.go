package service

import (
	"context"

	"github.com/rs401/letsgorip/forums/repository"
	"github.com/rs401/letsgorip/pb"
)

type forumService struct {
	forumsRepository repository.ForumsRepository
	pb.UnimplementedForumServiceServer
}

func NewForumService(forumsRepository repository.ForumsRepository) pb.ForumServiceServer {
	return &forumService{forumsRepository: forumsRepository}
}

func (fs *forumService) CreateForum(context.Context, *pb.Forum) (*pb.ForumIdResponse, error) {

}

func (fs *forumService) CreateThread(context.Context, *pb.Thread) (*pb.ForumIdResponse, error) {

}

func (fs *forumService) CreatePost(context.Context, *pb.Post) (*pb.ForumIdResponse, error) {

}

func (fs *forumService) GetForum(context.Context, *pb.ForumIdRequest) (*pb.Forum, error) {

}

func (fs *forumService) GetForums(*pb.GetForumsRequest, pb.ForumService_GetForumsServer) error {

}

func (fs *forumService) GetThread(context.Context, *pb.ForumIdRequest) (*pb.Thread, error) {

}

func (fs *forumService) GetThreads(*pb.GetThreadsRequest, pb.ForumService_GetThreadsServer) error {

}

func (fs *forumService) GetPosts(*pb.ForumIdRequest, pb.ForumService_GetPostsServer) error {

}

func (fs *forumService) UpdateForum(context.Context, *pb.Forum) (*pb.Forum, error) {

}

func (fs *forumService) UpdateThread(context.Context, *pb.Thread) (*pb.Thread, error) {

}

func (fs *forumService) UpdatePost(context.Context, *pb.Post) (*pb.Post, error) {

}

func (fs *forumService) DeleteForum(context.Context, *pb.ForumIdRequest) (*pb.ForumIdResponse, error) {

}

func (fs *forumService) DeleteThread(context.Context, *pb.ForumIdRequest) (*pb.ForumIdResponse, error) {

}

func (fs *forumService) DeletePost(context.Context, *pb.ForumIdRequest) (*pb.ForumIdResponse, error) {

}
