package models

import (
	"time"
)

type Model struct {
	ID int64 `gorm:"primaryKey;autoIncrement;comment:主键编码" json:"id"`
	ModelTime
}

type ControlBy struct {
	CreateBy int `json:"create_by" gorm:"index;comment:创建者"`
	UpdateBy int `json:"update_by" gorm:"index;comment:更新者"`
}

type ModelTime struct {
	CreateTime time.Time `gorm:"column:create_time; autoUpdateTime; default:CURRENT_TIMESTAMP;<-:create" json:"create_time,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time; autoCreateTime; default:CURRENT_TIMESTAMP on update current_timestamp" json:"update_time,omitempty"`
	//DeleteTime	gorm.DeleteAt 	`json:"-" gorm:"index;comment:删除时间"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}
