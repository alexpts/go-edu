package model

type Post struct {
	Model
	VersionMixin

	Title  string `gorm:"type:varchar(255)"`
	Status uint32
	UserId uint
	CatId  *string

	Category Category `gorm:"foreignKey:CatId"` // default FK CategoryID (relName + ID); references:id
}

func (m *Post) TableName() string {
	return `posts`
}

type PostRel struct {
	Post

	// Relations
	User     User
	Category Category
}
