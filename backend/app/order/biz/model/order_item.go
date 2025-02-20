package model

import (
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/2/16 21:03
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type OrderItem struct {
	gorm.Model
	ProductId     uint32  `gorm:"not null"`
	OrderParentId string  `gorm:"type:varchar(255) not null; index"`
	Quantity      uint32  `gorm:"not null"`
	Price         float32 `gorm:"not null; type:decimal(10,2)"`
}

func (o OrderItem) TableName() string {
	return "order_item"
}
