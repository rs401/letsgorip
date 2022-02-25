// Package handlers provides handlerfuncs
package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs401/letsgorip/api/tokenutils"
	"github.com/rs401/letsgorip/forums/models"
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
	var forum models.Forum
	// Decode body
	err := json.NewDecoder(r.Body).Decode(&forum)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	// Create
	response, err := fh.forumSvcClient.CreateForum(r.Context(), forum.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	forum.Id = response.GetId()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(forum)
}

func (fh *forumHandlers) CreateThread(w http.ResponseWriter, r *http.Request) {
	var thread models.Thread
	// Decode body
	err := json.NewDecoder(r.Body).Decode(&thread)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	// Create
	response, err := fh.forumSvcClient.CreateThread(r.Context(), thread.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	thread.Id = response.Id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(thread)
}

func (fh *forumHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	// Decode body
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	// Create
	response, err := fh.forumSvcClient.CreatePost(r.Context(), post.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	post.Id = response.Id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (fh *forumHandlers) GetForum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req pb.ForumIdRequest
	var forum models.Forum
	fid, err := strconv.Atoi(vars["fid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint64(fid)
	pbForum, err := fh.forumSvcClient.GetForum(r.Context(), &req)
	if err != nil {
		log.Printf("Error calling GetForum: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	forum.FromProtoBuffer(pbForum)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(forum)
}

func (fh *forumHandlers) GetThread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req pb.ForumIdRequest
	var thread models.Thread
	tid, err := strconv.Atoi(vars["tid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint64(tid)
	pbThread, err := fh.forumSvcClient.GetThread(r.Context(), &req)
	if err != nil {
		log.Printf("Error calling GetThread: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	thread.FromProtoBuffer(pbThread)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(thread)
}

func (fh *forumHandlers) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req pb.ForumIdRequest
	var post models.Post
	pid, err := strconv.Atoi(vars["pid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint64(pid)
	pbPost, err := fh.forumSvcClient.GetPost(r.Context(), &req)
	if err != nil {
		log.Printf("Error calling GetPost: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	post.FromProtoBuffer(pbPost)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (fh *forumHandlers) GetForums(w http.ResponseWriter, r *http.Request) {
	var gfRequest pb.GetForumsRequest
	var forums []*models.Forum = make([]*models.Forum, 0)
	var pbForums []*pb.Forum = make([]*pb.Forum, 0)

	done := make(chan bool)
	getForumsStream, err := fh.forumSvcClient.GetForums(r.Context(), &gfRequest)
	if err != nil {
		log.Printf("Error calling GetForums: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	go func() {
		for {
			response, err := getForumsStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Printf("Error receiving list forums stream")
				done <- false
				return
			}
			pbForums = append(pbForums, response)
		}
	}()

	if <-done {
		for _, f := range pbForums {
			forum := &models.Forum{}
			forum.FromProtoBuffer(f)
			forums = append(forums, forum)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(forums)
}

func (fh *forumHandlers) GetThreads(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req pb.ForumIdRequest
	var threads []*models.Thread = make([]*models.Thread, 0)
	var pbThreads []*pb.Thread = make([]*pb.Thread, 0)
	fid, err := strconv.Atoi(vars["fid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint64(fid)

	done := make(chan bool)
	getThreadsStream, err := fh.forumSvcClient.GetThreads(r.Context(), &req)
	if err != nil {
		log.Printf("Error calling GetThreads: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	go func() {
		for {
			response, err := getThreadsStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Printf("Error receiving list threads stream")
				done <- false
				return
			}
			pbThreads = append(pbThreads, response)
		}
	}()

	if <-done {
		for _, t := range pbThreads {
			thread := &models.Thread{}
			thread.FromProtoBuffer(t)
			threads = append(threads, thread)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(threads)
}

func (fh *forumHandlers) GetPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req pb.ForumIdRequest
	var posts []*models.Post = make([]*models.Post, 0)
	var pbPosts []*pb.Post = make([]*pb.Post, 0)
	tid, err := strconv.Atoi(vars["tid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint64(tid)

	done := make(chan bool)
	getPostsStream, err := fh.forumSvcClient.GetPosts(r.Context(), &req)
	if err != nil {
		log.Printf("Error calling GetPosts: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	go func() {
		for {
			response, err := getPostsStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Printf("Error receiving list threads stream")
				done <- false
				return
			}
			pbPosts = append(pbPosts, response)
		}
	}()

	if <-done {
		for _, p := range pbPosts {
			post := &models.Post{}
			post.FromProtoBuffer(p)
			posts = append(posts, post)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (fh *forumHandlers) UpdateForum(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	fid, err := strconv.Atoi(vars["fid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	var forum models.Forum
	err = json.NewDecoder(r.Body).Decode(&forum)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	if forum.Id != uint64(fid) {
		log.Printf("ForumId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	if forum.UserId != uint64(userId) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}

	result, err := fh.forumSvcClient.UpdateForum(r.Context(), forum.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	var tmpForum models.Forum
	tmpForum.FromProtoBuffer(result)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tmpForum)
}

func (fh *forumHandlers) UpdateThread(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	fid, err := strconv.Atoi(vars["fid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	tid, err := strconv.Atoi(vars["tid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	var thread models.Thread
	err = json.NewDecoder(r.Body).Decode(&thread)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	if thread.ForumId != uint64(fid) {
		log.Printf("ForumId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	if thread.Id != uint64(tid) {
		log.Printf("ThreadId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	if thread.UserId != uint64(userId) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}

	result, err := fh.forumSvcClient.UpdateThread(r.Context(), thread.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	var tmpThread models.Thread
	tmpThread.FromProtoBuffer(result)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tmpThread)
}

func (fh *forumHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	tid, err := strconv.Atoi(vars["tid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	pid, err := strconv.Atoi(vars["pid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	var post models.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	if post.ThreadId != uint64(tid) {
		log.Printf("ThreadId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	if post.Id != uint64(pid) {
		log.Printf("PostId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	if post.UserId != uint64(userId) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}

	result, err := fh.forumSvcClient.UpdatePost(r.Context(), post.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	var tmpPost models.Post
	tmpPost.FromProtoBuffer(result)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tmpPost)
}

func (fh *forumHandlers) DeleteForum(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	fid, err := strconv.Atoi(vars["fid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	pbForum, err := fh.forumSvcClient.GetForum(r.Context(), &pb.ForumIdRequest{Id: uint64(fid)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if pbForum.UserId != uint64(userId) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	result, err := fh.forumSvcClient.DeleteForum(r.Context(), &pb.ForumIdRequest{Id: uint64(fid)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if result.Id != uint64(fid) {
		// Uh oh
		log.Printf("Delete forum response.Id != path id. Got: %d, Expected: %d\n", result.Id, fid)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"delete": "success"})
}

func (fh *forumHandlers) DeleteThread(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	tid, err := strconv.Atoi(vars["tid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	pbThread, err := fh.forumSvcClient.GetThread(r.Context(), &pb.ForumIdRequest{Id: uint64(tid)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if pbThread.UserId != uint64(userId) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	result, err := fh.forumSvcClient.DeleteThread(r.Context(), &pb.ForumIdRequest{Id: uint64(tid)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if result.Id != uint64(tid) {
		// Uh oh
		log.Printf("Delete thread response.Id != path id. Got: %d, Expected: %d\n", result.Id, tid)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"delete": "success"})
}

func (fh *forumHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	pid, err := strconv.Atoi(vars["pid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	pbPost, err := fh.forumSvcClient.GetPost(r.Context(), &pb.ForumIdRequest{Id: uint64(pid)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if pbPost.UserId != uint64(userId) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	result, err := fh.forumSvcClient.DeletePost(r.Context(), &pb.ForumIdRequest{Id: uint64(pid)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if result.Id != uint64(pid) {
		// Uh oh
		log.Printf("Delete post response.Id != path id. Got: %d, Expected: %d\n", result.Id, pid)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"delete": "success"})
}
