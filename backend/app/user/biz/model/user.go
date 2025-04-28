package model

import (
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/2/3 22:32
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: 用户模型
 */

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex:idx_email;type:varchar(255) not null;"`
	PasswordHash string `gorm:"type:varchar(255) not null;"`
	Username     string `gorm:"type:varchar(255) not null;"`
	Avatar       string `gorm:"type:varchar(255) not null;"`
}

func (User) TableName() string {
	return "user"
}
