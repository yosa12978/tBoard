package models

const (
	ROLE_USER uint8 = iota
	ROLE_ADMIN
)

type Account struct {
	BaseModel
	Username string
	Password string    `json:"-"`
	Posts    []Post    `gorm:"foreignKey:AuthorId" json:",omitempty"`
	Comments []Comment `gorm:"foreignKey:AuthorId" json:",omitempty"`
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
