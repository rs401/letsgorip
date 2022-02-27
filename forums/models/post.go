// Package models provides the data models for the forum service.
package models

import "github.com/rs401/letsgorip/pb"

// Post is the data model for a Post.
type Post struct {
	Id        uint64 `json:"id" gorm:"primaryKey"`
	ThreadId  uint64 `json:"thread_id" gorm:"not null"`
	UserId    uint64 `json:"user_id" gorm:"not null"`
	Msg       string `json:"msg" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime"`
}

// ToProtoBuffer returns a protocol buffers version of the Post.
func (p *Post) ToProtoBuffer() *pb.Post {
	return &pb.Post{
		Id:        p.Id,
		ThreadId:  p.ThreadId,
		UserId:    p.UserId,
		Msg:       p.Msg,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// FromProtoBuffer takes a pb post and 'loads' the current post with the
// details.
func (p *Post) FromProtoBuffer(post *pb.Post) {
	p.Id = post.GetId()
	p.ThreadId = post.GetThreadId()
	p.UserId = post.GetUserId()
	p.Msg = post.GetMsg()
	p.CreatedAt = post.GetCreatedAt()
	p.UpdatedAt = post.GetUpdatedAt()
}
