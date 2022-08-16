package post

import (
	"gorm.io/gorm"
	"mendota-backend/internal/user"
)

type Post struct {
	gorm.Model
	Title   string
	Content string

	UserID uint `gorm:"index"`
	User   user.User
}
