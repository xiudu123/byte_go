package service

import (
	"byte_go/backend/app/order/biz/dal/mysql"
	"byte_go/backend/app/order/biz/model"
	order "byte_go/backend/rpc_gen/kitex_gen/order"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {

	// 校验参数
	if req == nil || req.UserId == 0 {
		return nil, kitex_err.RequestParamError
	}
	if len(req.OrderItems) == 0 {
		return nil, kitex_err.OrderItemEmpty
	}

	// 生成订单
	orderId, _ := uuid.NewRandom()
	o := &model.Order{
		OrderId:      orderId.String(),
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Consignee: model.Consignee{
			Email:    req.Email,
			Nickname: req.Nickname,
		},
	}
	if req.Address != nil {
		o.Consignee.StreetAddress = req.Address.StreetAddress
		o.Consignee.City = req.Address.City
		o.Consignee.State = req.Address.State
		o.Consignee.Country = req.Address.Country
		o.Consignee.ZipCode = req.Address.ZipCode
	}

	// 生成订单商品
	var items = make([]model.OrderItem, len(req.OrderItems))
	for i, item := range req.OrderItems {
		items[i] = model.OrderItemGen2Model(o.OrderId, item)
	}

	// 保存订单
	orderQuery := model.NewOrderQuery(s.ctx, mysql.DB)

	// 开启事务
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err = orderQuery.CreateOrder(o); err != nil {
			return err
		}
		if err = orderQuery.CreateItem(items); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		klog.Errorf("user of userId %d place order err: %v", req.UserId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 生成返回值
	return &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: o.OrderId,
		},
	}, nil
}
