package model

import (
	"gorm.io/plugin/optimisticlock"
	"time"
)

type Model struct {
	Id    uint      `gorm:"primaryKey; ->; <-:create"`
	CreAt time.Time `gorm:"->;<-:create; autoUpdateTime:milli; not null"`
	UpAt  time.Time `gorm:"autoUpdateTime:milli"`
}

type VersionMixin struct {
	Ver optimisticlock.Version `gorm:"column:_ver"`
	// Ver uint.Version `gorm:"column:_ver"`
}

//func (m *User) BeforeUpdate(tx *gorm.DB) (err error) {
//	dbname := "_ver"
//
//	stmt := tx.Statement
//	dv := reflect.ValueOf(stmt.Dest)
//
//	if reflect.Indirect(dv).Kind() == reflect.Struct {
//		selectColumns, restricted := stmt.SelectAndOmitColumns(false, true)
//
//		sd, _ := schema.Parse(stmt.Dest, &sync.Map{}, stmt.DB.NamingStrategy)
//		d := make(map[string]interface{})
//		for _, field := range sd.Fields {
//			//if field.DBName == dbname || field.DBName == stmt.Schema.PrioritizedPrimaryField.DBName {
//			if field.DBName == stmt.Schema.PrioritizedPrimaryField.DBName {
//				continue
//			}
//
//			if v, ok := selectColumns[field.DBName]; (ok && v) || (!ok && (!restricted || !stmt.SkipHooks)) {
//				if field.AutoUpdateTime > 0 {
//					continue
//				}
//
//				val, isZero := field.ValueOf(stmt.Context, dv)
//				if (ok || !isZero) && field.Updatable {
//					d[field.DBName] = val
//				}
//			}
//		}
//
//		stmt.Dest = d
//	}
//
//	stmt.SetColumn(dbname, gorm.Expr("_ver + ?", 1), true)
//	return tx.Error
//}
