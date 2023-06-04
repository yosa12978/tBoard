package models

type Post struct {
	BaseModel
	Content   string
	Timestamp uint64
	Comments  []Comment `gorm:"foreignKey:PostId" json:",omitempty"`
	AuthorId  uint
}

type PostCreateDTO struct {
	Content string
}
