package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"index;not null"`
	PostID  uint   `gorm:"index;not null"`
	User    User
	Post    Post
}
