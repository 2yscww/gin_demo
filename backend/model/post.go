package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserID     uint      `json:"user_id gorm:"not null"`
	CategoryID uint      `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50)"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp"`
}

func (post *Post) BeforeCreate(ctx *gorm.DB) (err error) {
	if post.ID == uuid.Nil {
		post.ID = uuid.New()
	}
	return
}
