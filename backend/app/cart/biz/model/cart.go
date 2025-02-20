package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/2/16 11:31
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"not null;index:idx_userid"`
	ProductId uint32 `gorm:"not null"`
	Quantity  uint32 `gorm:"not null"`
}

func (c Cart) TableName() string {
	return "cart"
}

type CartQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewCartQuery(ctx context.Context, db *gorm.DB) *CartQuery {
	return &CartQuery{
		ctx: ctx,
		db:  db,
	}
}

func (q CartQuery) AddItem(item *Cart) error {
	// 先查询是否存在
	var row Cart
	err := q.db.WithContext(q.ctx).Model(&Cart{}).
		Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 存在则更新
	if row.ID > 0 {
		return q.db.WithContext(q.ctx).Model(&Cart{}).
			Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			Update("quantity", gorm.Expr("quantity + ?", item.Quantity)).Error
	}

	// 不存在则创建
	return q.db.WithContext(q.ctx).Model(&Cart{}).Create(item).Error
}

func (q CartQuery) GetCartByUserId(userId uint32) (cart []*Cart, err error) {
	err = q.db.WithContext(q.ctx).Model(&Cart{}).
		Where(&Cart{UserId: userId}).
		Find(&cart).Error
	return
}

func (q CartQuery) EmptyCartByUserId(userId uint32) (err error) {
	err = q.db.WithContext(q.ctx).Model(&Cart{}).
		Where(&Cart{UserId: userId}).
		Delete(&Cart{}).Error
	return
}
