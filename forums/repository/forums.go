package repository

import (
	"errors"

	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/forums/models"
	"github.com/rs401/letsgorip/validation"
	"gorm.io/gorm"
)

var ErrorBadID error = errors.New("bad id")

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
}

type forumsRepository struct {
	db *gorm.DB
}

func NewForumsRepository(conn db.Connection) ForumsRepository {
	return &forumsRepository{db: conn.DB()}
}

func (fr *forumsRepository) CreateForum(forum *models.Forum) error {
	return fr.db.Create(&forum).Error
}

func (fr *forumsRepository) GetForum(id uint64) (forum *models.Forum, err error) {
	result := fr.db.Where("id = ?", id).First(&forum)
	return forum, result.Error
}

func (fr *forumsRepository) GetForumByTitle(title string) (forum *models.Forum, err error) {
	result := fr.db.Where("title = ?", title).First(&forum)
	return forum, result.Error
}

func (fr *forumsRepository) GetForums() (forums []*models.Forum, err error) {
	result := fr.db.Find(&forums)
	return forums, result.Error
}

func (fr *forumsRepository) UpdateForum(forum *models.Forum) error {
	var tmpForum = new(models.Forum)
	fr.db.Find(&tmpForum, forum.Id)
	if tmpForum.Title != forum.Title && !validation.IsEmptyString(forum.Title) {
		tmpForum.Title = forum.Title
	}
	if tmpForum.Description != forum.Description && validation.IsValidEmail(forum.Description) {
		tmpForum.Description = forum.Description
	}
	return fr.db.Save(&tmpForum).Error
}

func (fr *forumsRepository) DeleteForum(id uint64) error {
	var forum models.Forum
	fr.db.Find(&forum, id)
	if forum.Id == 0 {
		return ErrorBadID
	}
	return fr.db.Delete(&forum).Error
}
