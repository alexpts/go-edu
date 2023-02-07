package post

import (
	"time"
)

type Post struct {
	Id        uint `gorm:"primaryKey"`
	Title     string
	Status    uint
	CreatedAt time.Time
	UpdatedAt time.Time //`gorm:"autoCreateTime:false"`
}
