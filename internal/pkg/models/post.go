package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content   string
	Timestamp uint64
	Comments  []Comment `gorm:"foreignKey:PostId"`
	AuthorId  uint
}

type PostCreateDTO struct {
	Content string
}
