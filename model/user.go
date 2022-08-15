package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Password    string
	Salt        string
	HashedEmail string
	Avatar      string
	Bio         string
	Likes       uint
	DisLikes    uint
	Posts       []Post
}
