package models

type Forum struct {
	Id          uint64    `json:"id" gorm:"primaryKey"`
	UserId      uint64    `json:"user_id" gorm:"not null"`
	Title       string    `json:"title" gorm:"<-;unique;not null"`
	Description string    `json:"description" gorm:"not null"`
	Threads     []*Thread `json:"threads,omitempty"`
	CreatedAt   int64     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64     `json:"updated_at" gorm:"autoUpdateTime"`
}
