package model

type Post struct {
	Model

	Title  string `gorm:"type:varchar(255)"`
	Status uint32
	UserId uint
}

func (m *Post) TableName() string {
	return `posts`
}

type PostRel struct {
	Post
	// Rel
	User User
}
