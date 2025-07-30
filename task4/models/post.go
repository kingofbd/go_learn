package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"type:text;not null"`
	UserID   uint   `gorm:"index;not null"`
	User     User
	Comments []Comment
}
