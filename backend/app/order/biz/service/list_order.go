package service

import (
	"byte_go/backend/app/order/biz/dal/mysql"
	"byte_go/backend/app/order/biz/model"
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
	// 校验参数
	if req == nil || req.UserId == 0 {
		return nil, kitex_err.RequestParamError
	}

	// 查询订单
	orderQuery := model.NewOrderQuery(s.ctx, mysql.DB)
	orderList, err := orderQuery.ListOrder(req.UserId)
	if err != nil {
		klog.Errorf("user of userId %d list order err: %v", req.UserId, err)
		return nil, kitex_err.MysqlError
	}

	// 组装订单
	orders := make([]*order.Order, len(orderList))
	for i, o := range orderList {
		orders[i] = model.OrderModel2Gen(&o)
	}

	// 返回
	return &order.ListOrderResp{
		Orders: orders,
	}, nil
}
