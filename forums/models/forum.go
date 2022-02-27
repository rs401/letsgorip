// Package models provides the data models for the forum service.
package models

import "github.com/rs401/letsgorip/pb"

// Forum is the data model for a Forum.
type Forum struct {
	Id          uint64    `json:"id" gorm:"primaryKey"`
	UserId      uint64    `json:"user_id" gorm:"not null"`
	Title       string    `json:"title" gorm:"<-;unique;not null"`
	Description string    `json:"description" gorm:"not null"`
	Threads     []*Thread `json:"threads,omitempty"`
	CreatedAt   int64     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64     `json:"updated_at" gorm:"autoUpdateTime"`
}

// ToProtoBuffer returns a protocol buffers version of the Forum.
func (f *Forum) ToProtoBuffer() *pb.Forum {
	return &pb.Forum{
		Id:          f.Id,
		UserId:      f.UserId,
		Title:       f.Title,
		Description: f.Description,
		CreatedAt:   f.CreatedAt,
		UpdatedAt:   f.UpdatedAt,
	}
}

// FromProtoBuffer takes a pb forum and 'loads' the current forum with the
// details.
func (f *Forum) FromProtoBuffer(forum *pb.Forum) {
	f.Id = forum.GetId()
	f.UserId = forum.GetUserId()
	f.Title = forum.GetTitle()
	f.Description = forum.GetDescription()
	f.CreatedAt = forum.GetCreatedAt()
	f.UpdatedAt = forum.GetUpdatedAt()

}
