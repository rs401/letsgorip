package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs401/letsgorip/api/tokenutils"
	"github.com/rs401/letsgorip/pb"
	"github.com/rs401/letsgorip/places/models"
)

// PlaceHandlers is the interface defining the place handlerfuncs
type PlaceHandlers interface {
	CreatePlace(w http.ResponseWriter, r *http.Request)
	GetPlace(w http.ResponseWriter, r *http.Request)
	GetPlaces(w http.ResponseWriter, r *http.Request)
	UpdatePlace(w http.ResponseWriter, r *http.Request)
	DeletePlace(w http.ResponseWriter, r *http.Request)
}

type placeHandlers struct {
	placeSvcClient pb.PlaceServiceClient
}

// NewPlaceHandlers takes a pb.PlaceServiceClient and returns a PlaceHandlers
func NewPlaceHandlers(placeSvcClient pb.PlaceServiceClient) PlaceHandlers {
	return &placeHandlers{placeSvcClient: placeSvcClient}
}

// CreatePlace is the handlerfunc that makes the service client call to create a
// Place.
func (ph *placeHandlers) CreatePlace(w http.ResponseWriter, r *http.Request) {
	var place models.Place
	// Decode json place
	err := json.NewDecoder(r.Body).Decode(&place)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	// Create place
	response, err := ph.placeSvcClient.CreatePlace(r.Context(), place.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	place.Id = response.GetId()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(place)
}

// GetPlace is the handlerfunc that makes the service client call to retrieve a
// Place.
func (ph *placeHandlers) GetPlace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req pb.PlaceIdRequest
	var place models.Place
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	req.Id = uint64(id)
	pbPlace, err := ph.placeSvcClient.GetPlace(r.Context(), &req)
	if err != nil {
		log.Printf("Error calling GetPlace: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	place.FromProtoBuffer(pbPlace)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(place)
}

// GetPlaces is the handlerfunc that makes the service client call to retrieve
// all Places for a specific state.
func (ph *placeHandlers) GetPlaces(w http.ResponseWriter, r *http.Request) {
	var gpRequest pb.GetPlacesRequest
	var places []*models.Place = make([]*models.Place, 0)
	var pbPlaces []*pb.Place = make([]*pb.Place, 0)

	done := make(chan bool)
	getPlacesStream, err := ph.placeSvcClient.GetPlaces(r.Context(), &gpRequest)
	if err != nil {
		log.Printf("Error calling GetPlaces: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	go func() {
		for {
			response, err := getPlacesStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Printf("Error receiving places stream")
				done <- false
				return
			}
			pbPlaces = append(pbPlaces, response)
		}
	}()

	if <-done {
		for _, p := range pbPlaces {
			place := &models.Place{}
			place.FromProtoBuffer(p)
			places = append(places, place)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(places)
}

// UpdatePlace is the handlerfunc that makes the service client call update a
// Place.
func (ph *placeHandlers) UpdatePlace(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	var place models.Place
	err = json.NewDecoder(r.Body).Decode(&place)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json formating decode error"})
		return
	}
	if place.Id != uint64(id) {
		log.Printf("Place id did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	if place.UserId != uint64(userId) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}

	result, err := ph.placeSvcClient.UpdatePlace(r.Context(), place.ToProtoBuffer())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error from service client": err.Error()})
		return
	}
	var tmpPlace models.Place
	tmpPlace.FromProtoBuffer(result)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tmpPlace)
}

// DeletePlace is the handlerfunc that makes the service client call to delete a
// Place.
func (ph *placeHandlers) DeletePlace(w http.ResponseWriter, r *http.Request) {
	userId := tokenutils.ExtractUserId(r)
	if userId == 0 {
		log.Printf("UserId could not be extracted from token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized token error"})
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad path"})
		return
	}
	pbPlace, err := ph.placeSvcClient.GetPlace(r.Context(), &pb.PlaceIdRequest{Id: uint64(id)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if pbPlace.UserId != uint64(userId) {
		log.Printf("UserId did not match")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized id error"})
		return
	}
	result, err := ph.placeSvcClient.DeletePlace(r.Context(), &pb.PlaceIdRequest{Id: uint64(id)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if result.Id != uint64(id) {
		// Uh oh
		log.Printf("Delete place response.Id != path id. Got: %d, Expected: %d\n", result.Id, id)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"delete": "success"})
}
