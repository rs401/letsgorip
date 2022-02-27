// Package service implements the pb PlaceServiceServer
package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs401/letsgorip/pb"
	"github.com/rs401/letsgorip/places/models"
	"github.com/rs401/letsgorip/places/repository"
	"github.com/rs401/letsgorip/validation"
	"gorm.io/gorm"
)

type placeService struct {
	placesRepository repository.PlacesRepository
	pb.UnimplementedPlaceServiceServer
}

// NewPlaceService takes a places repository and returns a pb.PlaceServiceServer
func NewPlaceService(placesRepository repository.PlacesRepository) pb.PlaceServiceServer {
	return &placeService{placesRepository: placesRepository}
}

// CreatePlace takes a place and calls the repositories CreatePlace method.
func (ps *placeService) CreatePlace(ctx context.Context, req *pb.Place) (*pb.PlaceIdResponse, error) {
	// Check valid Place
	if err := validation.IsValidPlace(req); err != nil {
		return nil, err
	}
	// Check unique name if it exists
	exists, err := ps.placesRepository.GetPlaceByName(req.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			goto LabelContinue
		}
		return nil, err
	}
	if exists.Id != 0 {
		return nil, fmt.Errorf("error place name exists")
	}
LabelContinue:

	// Create Place
	place := new(models.Place)
	place.UserId = req.GetUserId()
	place.Name = req.GetName()
	place.Description = req.GetDescription()
	place.Latitude = req.GetLatitude()
	place.Longitude = req.GetLongitude()
	if err := ps.placesRepository.CreatePlace(place); err != nil {
		return nil, err
	}
	return &pb.PlaceIdResponse{Id: place.Id}, nil
}

// GetPlace takes a PlaceIdRequest and returns the corresponding Place.
func (ps *placeService) GetPlace(ctx context.Context, req *pb.PlaceIdRequest) (*pb.Place, error) {
	place, err := ps.placesRepository.GetPlace(req.GetId())
	if err != nil {
		return nil, err
	}
	return place.ToProtoBuffer(), nil
}

// GetPlaces takes a GetPlacesRequest and returns all places.
func (ps *placeService) GetPlaces(req *pb.GetPlacesRequest, stream pb.PlaceService_GetPlacesServer) error {
	places, err := ps.placesRepository.GetPlaces()
	if err != nil {
		return err
	}
	for _, place := range places {
		err := stream.Send(place.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdatePlace takes a place, validates it and calls the repositories
// UpdatePlace method.
func (ps *placeService) UpdatePlace(ctx context.Context, req *pb.Place) (*pb.Place, error) {
	// Verify place exists
	place, err := ps.placesRepository.GetPlace(req.GetId())
	if err != nil {
		return nil, err
	}
	if place == nil {
		return nil, validation.ErrNotFound
	}
	// Validate the place name not empty
	if validation.IsEmptyString(req.GetName()) {
		return nil, validation.ErrEmptyName
	}
	// Validate the place description not empty
	if validation.IsEmptyString(req.GetDescription()) {
		return nil, validation.ErrEmptyDescription
	}
	// Update record
	place.Name = req.GetName()
	place.Description = req.GetDescription()
	place.Latitude = req.GetLatitude()
	place.Longitude = req.GetLongitude()

	err = ps.placesRepository.UpdatePlace(place)
	return place.ToProtoBuffer(), err
}

// DeletePlace takes a PlaceIdRequest and calls the repositories DeletePlace
// method.
func (ps *placeService) DeletePlace(ctx context.Context, req *pb.PlaceIdRequest) (*pb.PlaceIdResponse, error) {
	err := ps.placesRepository.DeletePlace(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.PlaceIdResponse{Id: req.GetId()}, nil
}
