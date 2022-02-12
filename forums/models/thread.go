package models

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
