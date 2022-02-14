package repository

import (
	"github.com/rs401/letsgorip/forums/models"
	"github.com/rs401/letsgorip/validation"
)

func (fr *forumsRepository) CreateThread(thread *models.Thread) error {
	return fr.db.Create(&thread).Error
}

func (fr *forumsRepository) GetThread(id uint64) (thread *models.Thread, err error) {
	result := fr.db.Where("id = ?", id).First(&thread)
	return thread, result.Error
}

func (fr *forumsRepository) GetThreads(id uint64) (threads []*models.Thread, err error) {
	result := fr.db.Where("forum_id = ?", id).Find(&threads)
	return threads, result.Error
}

func (fr *forumsRepository) UpdateThread(thread *models.Thread) error {
	var tmpThread = new(models.Thread)
	fr.db.Find(&tmpThread, thread.Id)
	if tmpThread.Title != thread.Title && !validation.IsEmptyString(thread.Title) {
		tmpThread.Title = thread.Title
	}
	if tmpThread.Msg != thread.Msg && !validation.IsEmptyString(thread.Msg) {
		tmpThread.Msg = thread.Msg
	}
	return fr.db.Save(&tmpThread).Error
}

func (fr *forumsRepository) DeleteThread(id uint64) error {
	var thread models.Thread
	fr.db.Find(&thread, id)
	if thread.Id == 0 {
		return ErrorBadID
	}
	return fr.db.Delete(&thread).Error
}
