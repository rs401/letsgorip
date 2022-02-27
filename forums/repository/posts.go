// Package repository provides methods for interacting with the data store.
package repository

import (
	"github.com/rs401/letsgorip/forums/models"
	"github.com/rs401/letsgorip/validation"
)

// CreatePost takes a post and saves it to the database.
func (fr *forumsRepository) CreatePost(post *models.Post) error {
	return fr.db.Create(&post).Error
}

// GetPost takes an id and returns the post.
func (fr *forumsRepository) GetPost(id uint64) (post *models.Post, err error) {
	result := fr.db.Where("id = ?", id).First(&post)
	return post, result.Error
}

// GetPosts takes a thread id and returns all posts for that thread.
func (fr *forumsRepository) GetPosts(id uint64) (posts []*models.Post, err error) {
	result := fr.db.Where("thread_id = ?", id).Find(&posts)
	return posts, result.Error
}

// UpdatePost takes a post and updates it in the database.
func (fr *forumsRepository) UpdatePost(post *models.Post) error {
	var tmpPost = new(models.Post)
	fr.db.Find(&tmpPost, post.Id)
	if tmpPost.Msg != post.Msg && !validation.IsEmptyString(post.Msg) {
		tmpPost.Msg = post.Msg
	}
	return fr.db.Save(&tmpPost).Error
}

// DeletePost takes an id and deletes that post from the database.
func (fr *forumsRepository) DeletePost(id uint64) error {
	var post models.Post
	fr.db.Find(&post, id)
	if post.Id == 0 {
		return ErrorBadID
	}
	return fr.db.Delete(&post).Error
}

// DeleteAllPosts deletes all posts in the database.
func (fr *forumsRepository) DeleteAllPosts() error {
	return fr.db.Exec("DELETE FROM posts").Error
}
