// Package handlers provides handlerfuncs
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs401/letsgorip/api/tokenutils"
	"github.com/rs401/letsgorip/auth/models"
	"github.com/rs401/letsgorip/pb"
)

// AuthHandlers interface defining HandlerFuncs
type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	CheckToken(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {
	authSvcClient pb.AuthServiceClient
}

// NewAuthHandlers takes an authclient.AuthSvcClient and returns an AuthHandlers
func NewAuthHandlers(authSvcClient pb.AuthServiceClient) AuthHandlers {
	return &authHandlers{authSvcClient: authSvcClient}
}

// SignUp handles calling the Client.SignUp method
func (ah *authHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var pbUser pb.User
	err := json.NewDecoder(r.Body).Decode(&pbUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}

	result, err := ah.authSvcClient.SignUp(r.Context(), &pbUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Might as well create tokens so they don't have to log in
	tokens, err := tokenutils.CreateToken(uint(result.Id))
	if err != nil {
		log.Printf("Unable to create tokens after signup: %v\n", err)
	}
	if tokens != nil {
		// Set access token in auth header
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %v", tokens.AccessToken))
		// Set refresh token in cookie
		http.SetCookie(w,
			&http.Cookie{
				Name:     "refresh_token",
				Value:    tokens.RefreshToken,
				Expires:  time.Now().Add(time.Hour * 24),
				SameSite: http.SameSiteLaxMode,
			})
	}
	// Let them know
	user.FromProtoBuffer(&pbUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// SignIn handles calling the Client.SignIn method
func (ah *authHandlers) SignIn(w http.ResponseWriter, r *http.Request) {
	var signReq pb.SignInRequest
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&signReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	result, err := ah.authSvcClient.SignIn(r.Context(), &signReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if result.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}
	// All good, create tokens
	tokens, err := tokenutils.CreateToken(uint(result.Id))
	if err != nil {
		log.Printf("Unable to create tokens after signin: %v\n", err)
	}
	if tokens != nil {
		// Set access token in auth header
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %v", tokens.AccessToken))
		// Set refresh token in cookie
		http.SetCookie(w,
			&http.Cookie{
				Name:     "refresh_token",
				Value:    tokens.RefreshToken,
				Expires:  time.Now().Add(time.Hour * 24),
				SameSite: http.SameSiteLaxMode,
			})
	}
	// Let them know
	user.FromProtoBuffer(result)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// UpdateUser handles calling the Client.UpdateUser method
func (ah *authHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	var user = new(models.User)
	// var result = new(models.User)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	if userId != uint(id) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	if user.ID != uint64(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request, id error"})
		return
	}
	result, err := ah.authSvcClient.UpdateUser(r.Context(), user.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	var tmpUser models.User
	tmpUser.FromProtoBuffer(result)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tmpUser)
}

// GetUser handles calling the Client.GetUser method
func (ah *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req pb.GetUserRequest
	var user models.User
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint64(id)
	pbUser, err := ah.authSvcClient.GetUser(r.Context(), &req)
	if err != nil {
		log.Printf("Error calling GetUser: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	user.FromProtoBuffer(pbUser)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// GetUsers handles calling the Client.ListUsers method
func (ah *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	var luRequest pb.ListUsersRequest
	var users []*models.User = make([]*models.User, 0)
	var pbUsers []*pb.User = make([]*pb.User, 0)
	done := make(chan bool)
	listUsersStream, err := ah.authSvcClient.ListUsers(r.Context(), &luRequest)
	if err != nil {
		log.Printf("Error calling ListUsers: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	go func() {
		for {
			response, err := listUsersStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Printf("Error receiving list users stream")
				done <- false
				return
			}
			pbUsers = append(pbUsers, response)
		}
	}()

	if <-done {
		for _, u := range pbUsers {
			user := &models.User{}
			user.FromProtoBuffer(u)
			users = append(users, user)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// DeleteUser handles calling the Client.DeleteUser method
func (ah *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	// var req pb.GetUserRequest
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	if userId != uint(id) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}

	req := &pb.GetUserRequest{Id: uint64(id)}
	response, err := ah.authSvcClient.DeleteUser(r.Context(), req)
	if err != nil {
		log.Printf("Error calling DeleteUser: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if response.Id != uint64(id) {
		// Uh oh
		log.Printf("Delete user request response.Id != path id. Got: %d, Expected: %d\n", response.Id, id)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"delete": "success"})
}

func (ah *authHandlers) CheckToken(w http.ResponseWriter, r *http.Request) {
	token, err := tokenutils.VerifyToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if !tokenutils.StillValid(token) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"valid": "true"})
}
