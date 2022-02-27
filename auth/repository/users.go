// Package repository provides methods for interacting with the data store.
package repository

import (
	"errors"

	"github.com/rs401/letsgorip/auth/models"
	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/validation"
	"gorm.io/gorm"
)

// ErrorBadID is just an error indicating a 'bad id'.
var ErrorBadID error = errors.New("bad id")

// UsersRepository interface defines the method signatures.
type UsersRepository interface {
	Save(user *models.User) error
	GetById(id uint64) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Update(user *models.User) error
	Delete(id uint64) error
}

type usersRepository struct {
	db *gorm.DB
}

// NewUsersRepository takes a db connection and returns a UsersRepository.
func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{db: conn.DB()}
}

// Save takes a pointer to a models.User and saves the user to the database.
func (r *usersRepository) Save(user *models.User) error {
	return r.db.Create(&user).Error
}

// GetById takes an id and returns the user and an error.
func (r *usersRepository) GetById(id uint64) (user *models.User, err error) {
	result := r.db.Where("id = ?", id).First(&user)
	return user, result.Error
}

// GetByEmail takes an email string and returns the user and an error.
func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	result := r.db.Where("email = ?", email).Find(&user)
	return user, result.Error
}

// GetAll retrieves and returns all users.
func (r *usersRepository) GetAll() (users []*models.User, err error) {
	result := r.db.Find(&users)
	return users, result.Error
}

// Update takes a user and updates the user in the db.
func (r *usersRepository) Update(user *models.User) error {
	var tmpUser = new(models.User)
	r.db.Find(&tmpUser, user.ID)
	if tmpUser.Name != user.Name && !validation.IsEmptyString(user.Name) {
		tmpUser.Name = user.Name
	}
	if tmpUser.Email != user.Email && validation.IsValidEmail(user.Email) {
		tmpUser.Email = user.Email
	}
	return r.db.Save(&tmpUser).Error
}

// Delete takes an id and deletes the user from the db.
func (r *usersRepository) Delete(id uint64) error {
	var user models.User
	r.db.Find(&user, id)
	if user.ID == 0 {
		return ErrorBadID
	}
	return r.db.Delete(&user).Error
}

// DeleteAll deletes all users from the db.
func (r *usersRepository) DeleteAll() error {
	return r.db.Exec("DELETE FROM users").Error
}
