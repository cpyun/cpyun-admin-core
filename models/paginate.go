package models

import (
	"gorm.io/gorm"
)

func Paginate(pageSize, pageIndex int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if pageIndex <= 0 {
			pageIndex = 1
		}

		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (pageIndex - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}