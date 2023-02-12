package model

type User struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(50)"`

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
