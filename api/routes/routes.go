// Package routes provides utility to setup routes
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs401/letsgorip/api/handlers"
	"github.com/rs401/letsgorip/api/middlewares"
)

// SetupRoutes takes a *mux.Router and a AuthHandlers to configure *mux.Routes
func SetupRoutes(r *mux.Router, authHandleFuncs handlers.AuthHandlers) {
	r.HandleFunc("/api/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]string{"heartbeat": "alive"})
	})

	// Auth routes
	r.HandleFunc("/api/signup/", authHandleFuncs.SignUp).Methods("POST")
	r.HandleFunc("/api/signin/", authHandleFuncs.SignIn).Methods("POST")
	r.HandleFunc("/api/user/", authHandleFuncs.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/{id:[0-9]+}/", authHandleFuncs.GetUser).Methods("GET")

	// Forum routes
	// GET 		"/api/forum/"
	// GET 		"/api/forum/{fid:[0-9]+}/"

	// Thread routes
	// GET 		"/api/forum/{fid:[0-9]+}/thread/"
	// GET 		"/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/"

	// Post routes
	// GET 		"/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/"

	// Protected routes
	authRouter := r.PathPrefix("").Subrouter()
	authRouter.Use(middlewares.AuthMiddleware)
	authRouter.HandleFunc("/api/user/{id:[0-9]+}/", authHandleFuncs.UpdateUser).Methods("PUT")
	authRouter.HandleFunc("/api/user/{id:[0-9]+}/", authHandleFuncs.DeleteUser).Methods("DELETE")
	// Protected Forum routes
	// POST 	"/api/forum/"
	// PUT 		"/api/forum/{fid:[0-9]+}/"
	// DELETE 	"/api/forum/{fid:[0-9]+}/"
	// Protected Thread routes
	// POST 	"/api/forum/{fid:[0-9]+}/thread/"
	// PUT 		"/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/"
	// DELETE 	"/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/"
	// Protected Post routes
	// POST 		"/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/"
	// PUT 		"/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/{pid:[0-9]+}/"
	// DELETE 		"/api/forum/{fid:[0-9]+}/thread/{tid:[0-9]+}/post/{pid:[0-9]+}/"
}
