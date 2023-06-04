package models

type Comment struct {
	BaseModel
	Content  string
	PostId   uint
	AuthorId uint
}

type CommentCreateDTO struct {
	Content string
}
