package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content  string
	PostId   uint
	AuthorId uint
}

type CommentCreateDTO struct {
	Content string
}
