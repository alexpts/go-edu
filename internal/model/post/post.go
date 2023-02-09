package post

import (
	"github.com/alexpts/edu-go/internal/model/cat"
	"time"
)

type Post struct {
	Id        uint `gorm:"primaryKey"`
	Title     string
	Status    uint
	CreatedAt time.Time `gorm:"->"`
	UpdatedAt time.Time //`gorm:"autoCreateTime:false"`
	CatId     uint      `gorm:"foreignKey:cat_id;" json:"-"`
	// Relations
	Cat cat.Cat
}

// ShortPost alias model with custom read only short fields
type ShortPost struct {
	Id     uint   `gorm:"primaryKey; ->"`
	Title  string `gorm:"->"`
	Status uint   `gorm:"->"`
}

// TableName overrides the table name used by ShortPost to `posts`
func (ShortPost) TableName() string {
	return "posts"
}
