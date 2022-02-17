package repository

import (
	"github.com/rs401/letsgorip/forums/models"
	"github.com/rs401/letsgorip/validation"
)

func (fr *forumsRepository) CreatePost(post *models.Post) error {
	return fr.db.Create(&post).Error
}

func (fr *forumsRepository) GetPost(id uint64) (post *models.Post, err error) {
	result := fr.db.Where("id = ?", id).First(&post)
	return post, result.Error
}

func (fr *forumsRepository) GetPosts(id uint64) (posts []*models.Post, err error) {
	result := fr.db.Find(&posts)
	return posts, result.Error
}

func (fr *forumsRepository) UpdatePost(post *models.Post) error {
	var tmpPost = new(models.Post)
	fr.db.Find(&tmpPost, post.Id)
	if tmpPost.Msg != post.Msg && !validation.IsEmptyString(post.Msg) {
		tmpPost.Msg = post.Msg
	}
	return fr.db.Save(&tmpPost).Error
}

func (fr *forumsRepository) DeletePost(id uint64) error {
	var post models.Post
	fr.db.Find(&post, id)
	if post.Id == 0 {
		return ErrorBadID
	}
	return fr.db.Delete(&post).Error
}

func (fr *forumsRepository) DeleteAllPosts() error {
	return fr.db.Exec("DELETE FROM posts").Error
}
