// Package models provides the data models for the forum service.
package models

import "github.com/rs401/letsgorip/pb"

// Thread is the data model for a Thread.
type Thread struct {
	Id        uint64  `json:"id,omitempty" gorm:"primaryKey"`
	ForumId   uint64  `json:"forum_id" gorm:"not null"`
	UserId    uint64  `json:"user_id" gorm:"not null"`
	Title     string  `json:"title" gorm:"not null"`
	Msg       string  `json:"msg" gorm:"not null"`
	Posts     []*Post `json:"posts,omitempty"`
	CreatedAt int64   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64   `json:"updated_at" gorm:"autoUpdateTime"`
}

// ToProtoBuffer returns a protocol buffers version of the Thread.
func (t *Thread) ToProtoBuffer() *pb.Thread {
	return &pb.Thread{
		Id:        t.Id,
		ForumId:   t.ForumId,
		UserId:    t.UserId,
		Title:     t.Title,
		Msg:       t.Msg,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// FromProtoBuffer takes a pb thread and 'loads' the current thread with the
// details.
func (t *Thread) FromProtoBuffer(thread *pb.Thread) {
	t.Id = thread.GetId()
	t.ForumId = thread.GetForumId()
	t.UserId = thread.GetUserId()
	t.Title = thread.GetTitle()
	t.Msg = thread.GetMsg()
	t.CreatedAt = thread.GetCreatedAt()
	t.UpdatedAt = thread.GetUpdatedAt()
}
