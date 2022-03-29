// Package service implements the pb ForumServiceServer
package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs401/letsgorip/forums/models"
	"github.com/rs401/letsgorip/forums/repository"
	"github.com/rs401/letsgorip/pb"
	"github.com/rs401/letsgorip/validation"
	"gorm.io/gorm"
)

type forumService struct {
	forumsRepository repository.ForumsRepository
	pb.UnimplementedForumServiceServer
}

// NewForumService takes a forums repository and returns a pb.ForumServiceServer
func NewForumService(forumsRepository repository.ForumsRepository) pb.ForumServiceServer {
	return &forumService{forumsRepository: forumsRepository}
}

// CreateForum validates the forum and calls the repositories CreateForum method.
func (fs *forumService) CreateForum(ctx context.Context, req *pb.Forum) (*pb.ForumIdResponse, error) {
	// Check valid forum
	if err := validation.IsValidForum(req); err != nil {
		return nil, err
	}
	// Check if forum title exists
	exists, err := fs.forumsRepository.GetForumByTitle(req.Title)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			goto LabelContinue
		}
		return nil, err
	}
	if exists.Id != 0 {
		return nil, fmt.Errorf("error forum title exists")
	}
LabelContinue:

	// Create forum
	forum := new(models.Forum)
	forum.UserId = req.GetUserId()
	forum.Title = req.GetTitle()
	forum.Description = req.GetDescription()
	if err := fs.forumsRepository.CreateForum(forum); err != nil {
		return nil, err
	}
	return &pb.ForumIdResponse{Id: forum.Id}, nil
}

// CreateThread validates the thread and calls the repositories CreateThread method.
func (fs *forumService) CreateThread(ctx context.Context, req *pb.Thread) (*pb.ForumIdResponse, error) {
	// Check valid thread
	if err := validation.IsValidThread(req); err != nil {
		return nil, err
	}
	// Create thread
	thread := new(models.Thread)
	thread.ForumId = req.GetForumId()
	thread.UserId = req.GetUserId()
	thread.Title = req.GetTitle()
	thread.Msg = req.GetMsg()
	if err := fs.forumsRepository.CreateThread(thread); err != nil {
		return nil, err
	}
	return &pb.ForumIdResponse{Id: thread.Id}, nil
}

// CreatePost validates the post and calls the repositories CreatePost method.
func (fs *forumService) CreatePost(ctx context.Context, req *pb.Post) (*pb.ForumIdResponse, error) {
	// Check valid post
	if err := validation.IsValidPost(req); err != nil {
		return nil, err
	}
	// Create post
	post := new(models.Post)
	post.ThreadId = req.GetThreadId()
	post.UserId = req.GetUserId()
	post.Msg = req.GetMsg()
	if err := fs.forumsRepository.CreatePost(post); err != nil {
		return nil, err
	}
	return &pb.ForumIdResponse{Id: post.Id}, nil
}

// GetForum takes a ForumIdRequest and returns the Forum.
func (fs *forumService) GetForum(ctx context.Context, req *pb.ForumIdRequest) (*pb.Forum, error) {
	forum, err := fs.forumsRepository.GetForum(req.GetId())
	if err != nil {
		return nil, err
	}
	return forum.ToProtoBuffer(), nil
}

// GetForums takes a GetForumsRequest and returns all forums.
func (fs *forumService) GetForums(req *pb.GetForumsRequest, stream pb.ForumService_GetForumsServer) error {
	forums, err := fs.forumsRepository.GetForums()
	if err != nil {
		return err
	}
	for _, forum := range forums {
		err := stream.Send(forum.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}

// GetThread takes a ForumIdRequest and returns the Thread.
func (fs *forumService) GetThread(ctx context.Context, req *pb.ForumIdRequest) (*pb.Thread, error) {
	thread, err := fs.forumsRepository.GetThread(req.GetId())
	if err != nil {
		return nil, err
	}
	return thread.ToProtoBuffer(), nil
}

// GetThreads takes a ForumIdRequest and returns all threads for that forum.
func (fs *forumService) GetThreads(req *pb.ForumIdRequest, stream pb.ForumService_GetThreadsServer) error {
	threads, err := fs.forumsRepository.GetThreads(req.GetId())
	if err != nil {
		return err
	}
	for _, thread := range threads {
		err := stream.Send(thread.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}

// GetPost takes a ForumIdRequest and returns the Post.
func (fs *forumService) GetPost(ctx context.Context, req *pb.ForumIdRequest) (*pb.Post, error) {
	post, err := fs.forumsRepository.GetPost(req.GetId())
	if err != nil {
		return nil, err
	}
	return post.ToProtoBuffer(), nil
}

// GetPosts takes a ForumIdRequest(it will hold a thread id) and returns all
// posts for that thread.
func (fs *forumService) GetPosts(req *pb.ForumIdRequest, stream pb.ForumService_GetPostsServer) error {
	posts, err := fs.forumsRepository.GetPosts(req.GetId())
	if err != nil {
		return err
	}
	for _, post := range posts {
		err := stream.Send(post.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateForum takes a pb.Forum, validates the details and calls the
// repositories Update method.
func (fs *forumService) UpdateForum(ctx context.Context, req *pb.Forum) (*pb.Forum, error) {
	// Verify forum exists
	forum, err := fs.forumsRepository.GetForum(req.GetId())
	if err != nil {
		return nil, err
	}
	if forum == nil {
		return nil, validation.ErrNotFound
	}
	// Validate the forum title not empty
	if validation.IsEmptyString(req.GetTitle()) {
		return nil, validation.ErrEmptyTitle
	}
	// Validate the forum description not empty
	if validation.IsEmptyString(req.GetDescription()) {
		return nil, validation.ErrEmptyDescription
	}
	// Update record
	forum.Title = req.GetTitle()
	forum.Description = req.GetDescription()

	err = fs.forumsRepository.UpdateForum(forum)
	return forum.ToProtoBuffer(), err
}

// UpdateThread takes a pb.Thread, validates the details and calls the
// repositories Update method.
func (fs *forumService) UpdateThread(ctx context.Context, req *pb.Thread) (*pb.Thread, error) {
	// Verify thread exists
	thread, err := fs.forumsRepository.GetThread(req.GetId())
	if err != nil {
		return nil, err
	}
	if thread == nil {
		return nil, validation.ErrNotFound
	}
	// Validate the thread title not empty
	if validation.IsEmptyString(req.GetTitle()) {
		return nil, validation.ErrEmptyTitle
	}
	// Validate the thread message not empty
	if validation.IsEmptyString(req.GetMsg()) {
		return nil, validation.ErrEmptyMsg
	}
	// Update record
	thread.Title = req.GetTitle()
	thread.Msg = req.GetMsg()

	err = fs.forumsRepository.UpdateThread(thread)
	return thread.ToProtoBuffer(), err

}

// UpdatePost takes a pb.Post, validates the details and calls the repositories
// Update method.
func (fs *forumService) UpdatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	// Verify post exists
	post, err := fs.forumsRepository.GetPost(req.GetId())
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, validation.ErrNotFound
	}
	// Validate the post message not empty
	if validation.IsEmptyString(req.GetMsg()) {
		return nil, validation.ErrEmptyDescription
	}
	// Update record
	post.Msg = req.GetMsg()

	err = fs.forumsRepository.UpdatePost(post)
	return post.ToProtoBuffer(), err

}

// DeleteForum takes a ForumIdRequest and calls the repositories Delete method.
func (fs *forumService) DeleteForum(ctx context.Context, req *pb.ForumIdRequest) (*pb.ForumIdResponse, error) {
	err := fs.forumsRepository.DeleteForum(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ForumIdResponse{Id: req.GetId()}, nil
}

// DeleteThread takes a thread id in a ForumIdRequest and calls the repositories
// DeleteThread method.
func (fs *forumService) DeleteThread(ctx context.Context, req *pb.ForumIdRequest) (*pb.ForumIdResponse, error) {
	err := fs.forumsRepository.DeleteThread(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ForumIdResponse{Id: req.GetId()}, nil
}

// DeletePost takes a post id in a ForumIdRequest and calls the repositories
// DeletePost method.
func (fs *forumService) DeletePost(ctx context.Context, req *pb.ForumIdRequest) (*pb.ForumIdResponse, error) {
	err := fs.forumsRepository.DeletePost(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ForumIdResponse{Id: req.GetId()}, nil
}

func (fs *forumService) SearchForum(req *pb.ForumSearchRequest, stream pb.ForumService_SearchForumServer) error {
	threads, err := fs.forumsRepository.SearchForum(req.GetKey())
	if err != nil {
		return err
	}
	for _, thread := range threads {
		err := stream.Send(thread.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}
