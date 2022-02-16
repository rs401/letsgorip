package models

import "github.com/rs401/letsgorip/pb"

type Place struct {
	Id          uint64  `json:"id" gorm:"primaryKey"`
	UserId      uint64  `json:"user_id" gorm:"not null"`
	Name        string  `json:"name" gorm:"<-;unique;not null"`
	Description string  `json:"description" gorm:"not null"`
	Latitude    float64 `json:"latitude" gorm:"not null"`
	Longitude   float64 `json:"longitude" gorm:"not null"`
	CreatedAt   int64   `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt   int64   `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (p *Place) ToProtoBuffer() *pb.Place {
	return &pb.Place{
		Id:          p.Id,
		UserId:      p.UserId,
		Name:        p.Name,
		Description: p.Description,
		Latitude:    p.Latitude,
		Longitude:   p.Longitude,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p *Place) FromProtoBuffer(place *pb.Place) {
	p.Id = place.GetId()
	p.UserId = place.GetUserId()
	p.Name = place.GetName()
	p.Description = place.GetDescription()
	p.Latitude = place.GetLatitude()
	p.Longitude = place.GetLongitude()
	p.CreatedAt = place.GetCreatedAt()
	p.UpdatedAt = place.GetUpdatedAt()
}
