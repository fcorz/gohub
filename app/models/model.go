// Package models 模型通用属性和方法
package models

import (
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;comment:主键ID" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;comment:更新时间" json:"updated_at,omitempty"`
}
