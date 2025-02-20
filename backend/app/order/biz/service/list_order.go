package service

import (
	"byte_go/backend/app/order/biz/dal/mysql"
	"byte_go/backend/app/order/biz/model"
	"byte_go/backend/rpc_gen/kitex_gen/cart"
	order "byte_go/backend/rpc_gen/kitex_gen/order"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// 查询订单
	orderQuery := model.NewOrderQuery(s.ctx, mysql.DB)
	orderList, err := orderQuery.ListOrder(req.UserId)
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}

	// 组装订单
	orders := make([]*order.Order, len(orderList))
	for i, o := range orderList {

		// 组装订单商品
		orderItems := make([]*order.OrderItem, len(o.OrderItems))
		for j, item := range o.OrderItems {
			orderItems[j] = &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
				},
				Cost: item.Price,
			}
		}

		orders[i] = &order.Order{
			OrderId:      o.OrderId,
			UserId:       o.UserId,
			UserCurrency: o.UserCurrency,
			MarkedPaid:   o.MarkedPaid,
			Email:        o.Consignee.Email,
			CreatedAt:    int32(o.CreatedAt.Unix()),
			Address: &order.Address{
				StreetAddress: o.Consignee.StreetAddress,
				City:          o.Consignee.City,
				State:         o.Consignee.State,
				Country:       o.Consignee.Country,
				ZipCode:       o.Consignee.ZipCode,
			},
			OrderItems: orderItems,
		}

	}

	// 组装响应
	return &order.ListOrderResp{
		Orders: orders,
	}, nil
}
