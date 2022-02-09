// Package routes provides utility to setup routes
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs401/letsgorip/authapi/handlers"
	"github.com/rs401/letsgorip/authapi/middlewares"
)

// SetupRoutes takes a *mux.Router and a AuthHandlers to configure *mux.Routes
func SetupRoutes(r *mux.Router, hndFuncs handlers.AuthHandlers) {
	r.HandleFunc("/api/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]string{"heartbeat": "alive"})
	})

	r.HandleFunc("/api/signup/", hndFuncs.SignUp).Methods("POST")
	r.HandleFunc("/api/signin/", hndFuncs.SignIn).Methods("POST")
	r.HandleFunc("/api/user/", hndFuncs.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/{id:[0-9]+}/", hndFuncs.GetUser).Methods("GET")

	authRouter := r.PathPrefix("").Subrouter()
	authRouter.Use(middlewares.AuthMiddleware)
	authRouter.HandleFunc("/api/user/{id:[0-9]+}/", hndFuncs.UpdateUser).Methods("PUT")
	authRouter.HandleFunc("/api/user/{id:[0-9]+}/", hndFuncs.DeleteUser).Methods("DELETE")
}
