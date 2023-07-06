package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// FieldSortDest 自定义字段排序
func FieldSortDest(field string, bl bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: field,
			},
			Desc: bl,
		})
	}
}
