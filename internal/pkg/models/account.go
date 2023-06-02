package models

import "gorm.io/gorm"

const (
	ROLE_USER uint8 = iota
	ROLE_ADMIN
)

type Account struct {
	gorm.Model
	Username string
	Password string
	Posts    []Post    `gorm:"foreignKey:AuthorId"`
	Comments []Comment `gorm:"foreignKey:AuthorId"`
	Role     uint8
}

type LoginDTO struct {
	Username string
	Password string
}

type SignupDTO struct {
	Username string
	Password string
}

type UserInfo struct {
	Id       uint
	Username string
	Role     uint8
}
