// Package repository provides methods for interacting with the data store.
package repository

import (
	"github.com/rs401/letsgorip/forums/models"
	"github.com/rs401/letsgorip/validation"
)

// CreateThread takes a thread and saves it to the database.
func (fr *forumsRepository) CreateThread(thread *models.Thread) error {
	return fr.db.Create(&thread).Error
}

// GetThread takes an id and returns a thread.
func (fr *forumsRepository) GetThread(id uint64) (thread *models.Thread, err error) {
	result := fr.db.Where("id = ?", id).First(&thread)
	return thread, result.Error
}

// GetThreads takes a forum id and returns all threads for that forum.
func (fr *forumsRepository) GetThreads(id uint64) (threads []*models.Thread, err error) {
	result := fr.db.Where("forum_id = ?", id).Find(&threads)
	return threads, result.Error
}

// UpdateThread takes a thread and updates it in the database.
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

// DeleteThread takes an id and deletes the thread from the database.
func (fr *forumsRepository) DeleteThread(id uint64) error {
	var thread models.Thread
	fr.db.Find(&thread, id)
	if thread.Id == 0 {
		return ErrorBadID
	}
	return fr.db.Delete(&thread).Error
}

// DeleteAllThreads deletes all threads from the database.
func (fr *forumsRepository) DeleteAllThreads() error {
	return fr.db.Exec("DELETE FROM threads").Error
}
