// Package models provides the data models for the auth service
package models

import (
	"time"

	"github.com/rs401/letsgorip/pb"
)

// User model for user
type User struct {
	ID            uint64    `json:"id" gorm:"primaryKey"`
	Uid           string    `json:"uid" gorm:"<-;unique;not null"`
	Name          string    `json:"name" gorm:"<-;unique;not null"`
	Email         string    `json:"email" gorm:"<-;unique;not null"`
	EmailVerified bool      `json:"email_verified" `
	Picture       string    `json:"picture" `
	Role          int       `json:"-"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// ToProtoBuffer returns a protocol buffers version of the User.
func (u *User) ToProtoBuffer() *pb.User {
	return &pb.User{
		Id:            uint64(u.ID),
		Uid:           u.Uid,
		Name:          u.Name,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		Picture:       u.Picture,
		Role:          pb.UserRole(u.Role),
		CreatedAt:     u.CreatedAt.Unix(),
		UpdatedAt:     u.UpdatedAt.Unix(),
	}
}

// FromProtoBuffer takes a pb user and 'loads' the current user with the details
func (u *User) FromProtoBuffer(user *pb.User) {
	u.ID = user.GetId()
	u.Uid = user.GetUid()
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.EmailVerified = user.GetEmailVerified()
	u.Picture = user.GetPicture()
	u.Role = user.GetRole().Descriptor().Index()
	u.CreatedAt = time.Unix(user.CreatedAt, 0)
	u.UpdatedAt = time.Unix(user.UpdatedAt, 0)
}
