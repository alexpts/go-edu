package model

type User struct {
	Model
	Id   uint   `gorm:"primaryKey; ->; <-:create;"`
	Name string `gorm:"type:varchar(50)"`

	// Relations
	Posts []Post `gorm:"foreignKey:UserId"`
}

func (m *User) TableName() string {
	return `users`
}

type Profile struct {
	User
	// Rel
	Posts []Post `gorm:"foreignKey:UserId"`
}
