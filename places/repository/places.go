package repository

import (
	"errors"

	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/places/models"
	"github.com/rs401/letsgorip/validation"
	"gorm.io/gorm"
)

var ErrorBadID error = errors.New("bad id")

type PlacesRepository interface {
	CreatePlace(place *models.Place) error
	GetPlace(id uint64) (*models.Place, error)
	GetPlaceByName(name string) (*models.Place, error)
	GetPlaces() ([]*models.Place, error)
	UpdatePlace(place *models.Place) error
	DeletePlace(id uint64) error
}

type placesRepository struct {
	db *gorm.DB
}

func NewPlacesRepository(conn db.Connection) PlacesRepository {
	return &placesRepository{db: conn.DB()}
}

func (pr *placesRepository) CreatePlace(place *models.Place) error {
	return pr.db.Create(&place).Error
}

func (pr *placesRepository) GetPlace(id uint64) (place *models.Place, err error) {
	result := pr.db.Where("id = ?", id).First(&place)
	return place, result.Error
}

func (pr *placesRepository) GetPlaceByName(name string) (place *models.Place, err error) {
	result := pr.db.Where("name = ?", name).First(&place)
	return place, result.Error
}

func (pr *placesRepository) GetPlaces() (places []*models.Place, err error) {
	result := pr.db.Find(&places)
	return places, result.Error
}

func (pr *placesRepository) UpdatePlace(place *models.Place) error {
	var tmpPlace = new(models.Place)
	pr.db.Find(&tmpPlace, place.Id)
	if tmpPlace.Name != place.Name && !validation.IsEmptyString(place.Name) {
		tmpPlace.Name = place.Name
	}
	if tmpPlace.Description != place.Description && !validation.IsEmptyString(place.Description) {
		tmpPlace.Description = place.Description
	}
	return pr.db.Save(&tmpPlace).Error
}

func (pr *placesRepository) DeletePlace(id uint64) error {
	var place models.Place
	pr.db.Find(&place, id)
	if place.Id == 0 {
		return ErrorBadID
	}
	return pr.db.Delete(&place).Error
}
