package model

type User struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(50)"`

	// Rel
	Posts []Post
}

func (m *User) TableName() string {
	return `users`
}
