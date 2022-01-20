// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"not null;default:'';comment:用户姓名" json:"name,omitempty"`
	Email    string `gorm:"not null;default:'';comment:邮箱" json:"-"`
	Phone    string `gorm:"not null;default:'';comment:电话" json:"-"`
	Password string `gorm:"not null;default:'';comment:密码" json:"-"`

	models.CommonTimestampsField
}
