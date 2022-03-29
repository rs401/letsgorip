// Package repository provides methods for interacting with the data store.
package repository

import (
	"errors"
	"strings"

	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/forums/models"
	"github.com/rs401/letsgorip/validation"
	"gorm.io/gorm"
)

// ErrorBadID is an error indicating a bad id.
var ErrorBadID error = errors.New("bad id")

// ForumsRepository interface defines the method signatures.
type ForumsRepository interface {
	CreateForum(forum *models.Forum) error
	CreateThread(thread *models.Thread) error
	CreatePost(post *models.Post) error
	GetForum(id uint64) (*models.Forum, error)
	GetForumByTitle(title string) (*models.Forum, error)
	GetForums() ([]*models.Forum, error)
	GetThread(id uint64) (*models.Thread, error)
	GetThreads(id uint64) ([]*models.Thread, error)
	GetPost(id uint64) (*models.Post, error)
	GetPosts(id uint64) ([]*models.Post, error)
	UpdateForum(forum *models.Forum) error
	UpdateThread(thread *models.Thread) error
	UpdatePost(post *models.Post) error
	DeleteForum(id uint64) error
	DeleteThread(id uint64) error
	DeletePost(id uint64) error
	SearchForum(key string) ([]*models.Thread, error)
}

type forumsRepository struct {
	db *gorm.DB
}

// NewForumsRepository takes a db connection and returns a ForumsRepository.
func NewForumsRepository(conn db.Connection) ForumsRepository {
	return &forumsRepository{db: conn.DB()}
}

// CreateForum takes a forum and saves it to the database.
func (fr *forumsRepository) CreateForum(forum *models.Forum) error {
	return fr.db.Create(&forum).Error
}

// GetForum takes an id and retrieves and returns the Forum.
func (fr *forumsRepository) GetForum(id uint64) (forum *models.Forum, err error) {
	result := fr.db.Where("id = ?", id).First(&forum)
	return forum, result.Error
}

// GetForumByTitle takes a title string and retrieves and returns the forum.
func (fr *forumsRepository) GetForumByTitle(title string) (forum *models.Forum, err error) {
	result := fr.db.Where("title = ?", title).First(&forum)
	return forum, result.Error
}

// GetForums retrieves and returns all forums.
func (fr *forumsRepository) GetForums() (forums []*models.Forum, err error) {
	result := fr.db.Find(&forums)
	return forums, result.Error
}

// UpdateForum takes a forum and updates it in the db.
func (fr *forumsRepository) UpdateForum(forum *models.Forum) error {
	var tmpForum = new(models.Forum)
	fr.db.Find(&tmpForum, forum.Id)
	if tmpForum.Title != forum.Title && !validation.IsEmptyString(forum.Title) {
		tmpForum.Title = forum.Title
	}
	if tmpForum.Description != forum.Description && !validation.IsEmptyString(forum.Description) {
		tmpForum.Description = forum.Description
	}
	return fr.db.Save(&tmpForum).Error
}

// DeleteForum takes a forum id and deletes the forum from the database.
func (fr *forumsRepository) DeleteForum(id uint64) error {
	var forum models.Forum
	fr.db.Find(&forum, id)
	if forum.Id == 0 {
		return ErrorBadID
	}
	return fr.db.Delete(&forum).Error
}

// DeleteAllForums deletes all forums from the database.
func (fr *forumsRepository) DeleteAllForums() error {
	return fr.db.Exec("DELETE FROM forums").Error
}

func (fr *forumsRepository) SearchForum(key string) (threads []*models.Thread, err error) {
	var tmpThreads []*models.Thread
	result := fr.db.Find(&tmpThreads)
	for _, t := range tmpThreads {
		if strings.Contains(t.Title, key) || strings.Contains(t.Msg, key) {
			threads = append(threads, t)
		}
	}
	return threads, result.Error
}
