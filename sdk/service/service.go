package service

import (
	"fmt"
	"github.com/cpyun/cpyun-admin-core/logger"
	"gorm.io/gorm"
)

type Service struct {
	Orm       *gorm.DB
	Log       *logger.Helper
	ErrorCode string
	Error     error
}

func (db *Service) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}
