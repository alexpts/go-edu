package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id    uint      `gorm:"primaryKey; ->; <-:create"`
	CreAt time.Time `gorm:"->;<-:create; autoUpdateTime:milli"`
	UpAt  time.Time `gorm:"autoUpdateTime:milli"`
	Ver   uint      `gorm:"column:_ver"`
}

// BeforeUpdate update version of row for each update @todo question https://github.com/go-gorm/gorm/issues/6079
func (m *User) BeforeUpdate(tx *gorm.DB) (err error) {
	tx = tx.Set("_ver", gorm.Expr("_ver + ?", 1))
	//tx = tx.Table(tx.Statement.Table)
	//tx = tx.UpdateColumn("_ver", gorm.Expr("_ver + ?", 2))
	//tx = tx.Table("users").UpdateColumn("_ver", gorm.Expr("_ver + ?", 1))
	return tx.Error
}
