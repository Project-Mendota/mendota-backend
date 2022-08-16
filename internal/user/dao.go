package user

import (
	"gorm.io/gorm"
)

type UserStatus int
type UserRole int

const (
	Active UserStatus = iota
	Blocked
)

const (
	UserSimple UserRole = iota
	Admin
)

type User struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Password    string
	Salt        string
	HashedEmail string `gorm:"unique"`

	Status UserStatus `gorm:"default:0"`
	Role   UserRole   `gorm:"default:0"`
}

type Profile struct {
	gorm.Model
	Avatar   string
	Bio      string
	Likes    uint `gorm:"default:0"`
	DisLikes uint `gorm:"default:0"`

	UserID uint `gorm:"index"`
	User   User
}
