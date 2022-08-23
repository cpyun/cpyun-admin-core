package models

import (
	"time"
)

type Model struct {
	ID int64 `gorm:"primaryKey;autoIncrement;comment:主键编码" json:"id"`
	ModelTime
}

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

type ModelTime struct {
	CreatedTime time.Time `gorm:"column:created_time; autoUpdateTime; default:CURRENT_TIMESTAMP;<-:create" json:"created_time,omitempty"`
	UpdatedTime time.Time `gorm:"column:update_time; autoCreateTime; default:CURRENT_TIMESTAMP on update current_timestamp" json:"updated_time,omitempty"`
	//DeletedTime	gorm.DeletedAt 	`json:"-" gorm:"index;comment:删除时间"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}
