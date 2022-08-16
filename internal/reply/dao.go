package reply

import (
	"gorm.io/gorm"
	"mendota-backend/internal/user"
)

type Reply struct {
	gorm.Model

	UserID uint `gorm:"index"`
	User   user.User
}
