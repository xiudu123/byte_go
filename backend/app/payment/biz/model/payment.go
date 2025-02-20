package model

import (
	"context"
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/2/19 14:59
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type PaymentLog struct {
	gorm.Model
	UserId        int32   `gorm:"type:varchar(255);not null"`
	OrderId       string  `gorm:"type:varchar(255);not null"`
	TransactionId string  `gorm:"type:varchar(255);not null; uniqueIndex"`
	Amount        float32 `gorm:"type:decimal(10,2);not null"`
}

func (p PaymentLog) TableName() string {
	return "payment_log"
}

type PaymentQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewPaymentQuery(ctx context.Context, db *gorm.DB) *PaymentQuery {
	return &PaymentQuery{
		ctx: ctx,
		db:  db,
	}
}

func (q PaymentQuery) CreatePaymentLog(paymentLog *PaymentLog) (err error) {
	return q.db.WithContext(q.ctx).Model(&PaymentLog{}).Create(paymentLog).Error
}
