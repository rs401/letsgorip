package models

import (
	"time"

	"github.com/rs401/letsgorip/pb"
)

// User model for user
type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"<-;unique;not null"`
	Email     string    `json:"email" gorm:"<-;unique;not null"`
	Password  []byte    `json:"-"`
	Role      int       `json:"-"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *User) ToProtoBuffer() *pb.User {
	return &pb.User{
		Id:        uint64(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		Role:      pb.UserRole(u.Role),
		Password:  string(u.Password),
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}
}

func (u *User) FromProtoBuffer(user *pb.User) {
	u.ID = user.GetId()
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.Password = []byte(user.GetPassword())
	u.Role = user.GetRole().Descriptor().Index()
	u.CreatedAt = time.Unix(user.CreatedAt, 0)
	u.UpdatedAt = time.Unix(user.UpdatedAt, 0)
}
