package model

import (
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/2/14 22:25
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`

	Products []Product `gorm:"many2many:product_category;"`
}

func (c Category) TableName() string {
	return "category"
}
