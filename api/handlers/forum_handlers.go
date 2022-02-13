// Package handlers provides handlerfuncs
package handlers

import (
	"net/http"

	"github.com/rs401/letsgorip/pb"
)

// ForumHandlers interface defining HandlerFuncs
type ForumHandlers interface {
	CreateForum(w http.ResponseWriter, r *http.Request)
	CreateThread(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetForum(w http.ResponseWriter, r *http.Request)
	GetThread(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
	GetForums(w http.ResponseWriter, r *http.Request)
	GetThreads(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
	UpdateForum(w http.ResponseWriter, r *http.Request)
	UpdateThread(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeleteForum(w http.ResponseWriter, r *http.Request)
	DeleteThread(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type forumHandlers struct {
	forumSvcClient pb.ForumServiceClient
}

// NewForumHandlers takes a pb.ForumSvcClient and returns a ForumHandlers
func NewForumHandlers(forumSvcClient pb.ForumServiceClient) ForumHandlers {
	return &forumHandlers{forumSvcClient: forumSvcClient}
}

func (fh *forumHandlers) CreateForum(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) CreateThread(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) GetForum(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) GetThread(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) GetPost(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) GetForums(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) GetThreads(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) GetPosts(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) UpdateForum(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) UpdateThread(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) DeleteForum(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) DeleteThread(w http.ResponseWriter, r *http.Request) {

}

func (fh *forumHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {

}
