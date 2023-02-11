package model

import (
	"time"
)

type Post struct {
	Id        uint   `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(255)"`
	Status    uint32
	CreatedAt time.Time `gorm:"->" json:"CreAt"`
	UpdatedAt time.Time `json:"UpAt"` //`gorm:"autoCreateTime:false"`
	UserId    uint

	// Rel
	User User
}

func (m *Post) TableName() string {
	return `posts`
}
