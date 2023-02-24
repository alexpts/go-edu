package model

type User struct {
	Model
	VersionMixin

	Name  string `gorm:"type:varchar(50)"`
	Posts []Post `gorm:"foreignKey:UserId"`
}

// TableName all child (embed)
func (m *User) TableName() string {
	return `users`
}

//type Profile struct {
//	User
//	// Relations
//	Posts []Post `gorm:"foreignKey:UserId"`
//}
