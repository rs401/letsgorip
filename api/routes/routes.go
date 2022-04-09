// Package routes provides utility to setup routes
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs401/letsgorip/api/handlers"
	"github.com/rs401/letsgorip/api/middlewares"
)

// SetupRoutes takes a *mux.Router and various handlers to configure *mux.Routes
func SetupRoutes(r *mux.Router, authHandleFuncs handlers.AuthHandlers, forumHandleFuncs handlers.ForumHandlers, placeHandleFuncs handlers.PlaceHandlers) {
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]string{"healthcheck": "healthy"})
	})
	r.HandleFunc("/api/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]string{"heartbeat": "alive"})
	})

	// Auth routes
	// r.HandleFunc("/api/signup/", authHandleFuncs.SignUp).Methods("POST")
	r.HandleFunc("/api/signin/", authHandleFuncs.SignIn).Methods("POST")
	r.HandleFunc("/api/user/", authHandleFuncs.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/checktoken", authHandleFuncs.CheckToken).Methods("GET")
	r.HandleFunc("/api/user/{id:[0-9]+}/", authHandleFuncs.GetUser).Methods("GET")

	// Forum routes
	r.HandleFunc("/api/forum/", forumHandleFuncs.GetForums).Methods("GET")
	r.HandleFunc("/api/forum/{fid:[0-9]+}/", forumHandleFuncs.GetForum).Methods("GET")
	r.HandleFunc("/api/forum/{query}/", forumHandleFuncs.SearchForums).Methods("GET")
	// Thread routes
	r.HandleFunc("/api/forum/{fid:[0-9]+}/thread/", forumHandleFuncs.GetThreads).Methods("GET")
	r.HandleFunc("/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/", forumHandleFuncs.GetThread).Methods("GET")
	// Post routes
	r.HandleFunc("/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/", forumHandleFuncs.GetPosts).Methods("GET")
	r.HandleFunc("/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/{pid:[0-9]+}/", forumHandleFuncs.GetPost).Methods("GET")

	// Place routes
	// GET 		"/api/place/"
	r.HandleFunc("/api/place/", placeHandleFuncs.GetPlaces).Methods("GET")
	r.HandleFunc("/api/place/{id:[0-9]+}/", placeHandleFuncs.GetPlace).Methods("GET")

	// Protected routes
	authRouter := r.PathPrefix("").Subrouter()
	authRouter.Use(middlewares.AuthMiddleware)
	authRouter.HandleFunc("/api/user/{id:[0-9]+}/", authHandleFuncs.UpdateUser).Methods("PUT")
	authRouter.HandleFunc("/api/user/{id:[0-9]+}/", authHandleFuncs.DeleteUser).Methods("DELETE")

	// Protected Forum routes
	authRouter.HandleFunc("/api/forum/", forumHandleFuncs.CreateForum).Methods("POST")
	authRouter.HandleFunc("/api/forum/{fid:[0-9]+}/", forumHandleFuncs.UpdateForum).Methods("PUT")
	authRouter.HandleFunc("/api/forum/{fid:[0-9]+}/", forumHandleFuncs.DeleteForum).Methods("DELETE")
	// Protected Thread routes
	authRouter.HandleFunc("/api/forum/{fid:[0-9]+}/thread/", forumHandleFuncs.CreateThread).Methods("POST")
	authRouter.HandleFunc("/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/", forumHandleFuncs.UpdateThread).Methods("PUT")
	authRouter.HandleFunc("/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/", forumHandleFuncs.DeleteThread).Methods("DELETE")
	// Protected Post routes
	authRouter.HandleFunc("/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/", forumHandleFuncs.CreatePost).Methods("POST")
	authRouter.HandleFunc("/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/{pid:[0-9]+}/", forumHandleFuncs.UpdatePost).Methods("PUT")
	authRouter.HandleFunc("/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/{pid:[0-9]+}/", forumHandleFuncs.DeletePost).Methods("DELETE")
	// Protected Place routes
	authRouter.HandleFunc("/api/place/", placeHandleFuncs.CreatePlace).Methods("POST")
	authRouter.HandleFunc("/api/place/{id:[0-9]+}/", placeHandleFuncs.UpdatePlace).Methods("PUT")
	authRouter.HandleFunc("/api/place/{id:[0-9]+}/", placeHandleFuncs.DeletePlace).Methods("DELETE")
}
