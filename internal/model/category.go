package model

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

type Category struct {
	Model

	Id    string `gorm:"primaryKey; ->; <-:create; type:varchar(50)"` // uuid v4
	Title string `gorm:"type:varchar(50)"`
}

func (m *Category) TableName() string {
	return `cats`
}

func (m *Category) BeforeCreate(tx *gorm.DB) error {
	uuidV4, err := uuid.NewV4() // replace go uuid v7
	if err == nil {
		m.Id = uuidV4.String()
	}

	return fmt.Errorf("can`t make uuid: %w", err)
}
