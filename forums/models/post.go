package models

type Post struct {
	Id        uint64 `json:"id" gorm:"primaryKey"`
	ThreadId  uint64 `json:"thread_id" gorm:"not null"`
	UserId    uint64 `json:"user_id" gorm:"not null"`
	Msg       string `json:"msg" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime"`
}
