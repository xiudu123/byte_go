package model

import (
	"byte_go/backend/rpc_gen/kitex_gen/cart"
	"byte_go/backend/rpc_gen/kitex_gen/order"
	"context"
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/2/16 20:46
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type Consignee struct {
	Nickname      string
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	gorm.Model
	UserId       uint32      `gorm:"not null"`
	OrderId      string      `gorm:"type:varchar(255) not null; uniqueIndex"`
	UserCurrency string      `gorm:"type:varchar(255) not null"`
	MarkedPaid   bool        `gorm:"default:false"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderParentId; references:OrderId"`
}

func (o Order) TableName() string {
	return "order"
}

func OrderModel2Gen(orderModel *Order) (orderGen *order.Order) {
	orderItems := make([]*order.OrderItem, len(orderModel.OrderItems))
	for j, item := range orderModel.OrderItems {
		orderItems[j] = &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
			},
			Cost: item.Price,
		}
	}

	orderGen = &order.Order{
		OrderId:      orderModel.OrderId,
		UserId:       orderModel.UserId,
		UserCurrency: orderModel.UserCurrency,
		MarkedPaid:   orderModel.MarkedPaid,
		Email:        orderModel.Consignee.Email,
		CreatedAt:    int32(orderModel.CreatedAt.Unix()),
		Address: &order.Address{
			StreetAddress: orderModel.Consignee.StreetAddress,
			City:          orderModel.Consignee.City,
			State:         orderModel.Consignee.State,
			Country:       orderModel.Consignee.Country,
			ZipCode:       orderModel.Consignee.ZipCode,
		},
		OrderItems: orderItems,
	}

	return orderGen
}

type OrderQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewOrderQuery(ctx context.Context, db *gorm.DB) *OrderQuery {
	return &OrderQuery{
		ctx: ctx,
		db:  db,
	}
}

func (q OrderQuery) CreateOrder(order *Order) (err error) {
	return q.db.WithContext(q.ctx).Create(order).Error
}

func (q OrderQuery) CreateItem(items []OrderItem) (err error) {
	return q.db.WithContext(q.ctx).Create(items).Error
}

func (q OrderQuery) GetOrderByOrderId(orderId string) (order Order, err error) {
	err = q.db.WithContext(q.ctx).
		Where("order_id =?", orderId).
		Preload("OrderItems").
		First(&order).Error
	return order, err
}

func (q OrderQuery) ListOrder(userId uint32) (orders []Order, err error) {
	err = q.db.WithContext(q.ctx).
		Where("user_id =?", userId).
		Preload("OrderItems").
		Find(&orders).Error
	return orders, err
}

func (q OrderQuery) MarkOrderPaid(orderId string) (err error) {
	return q.db.WithContext(q.ctx).
		Model(&Order{}).
		Where("order_id =?", orderId).
		Update("marked_paid", true).Error
}
