package service

import (
	"byte_go/backend/app/order/biz/dal/mysql"
	"byte_go/backend/app/order/biz/model"
	order "byte_go/backend/rpc_gen/kitex_gen/order"
	"byte_go/kitex_err"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {

	// 参数校验
	if req == nil || req.OrderId == "" {
		return nil, kitex_err.RequestParamError
	}

	orderQuery := model.NewOrderQuery(s.ctx, mysql.DB)

	// 获取订单
	orderInfo, err := orderQuery.GetOrderByOrderId(req.OrderId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		klog.Errorf("order of orderId %s not exist", req.OrderId)
		return nil, kitex_err.OrderNotExist
	}
	if err != nil {
		klog.Errorf("order of orderId %s get err: %v", req.OrderId, err)
		return nil, kitex_err.MysqlError
	}
	if orderInfo.MarkedPaid {
		klog.Errorf("order of orderId %s has been paid", req.OrderId)
		return nil, kitex_err.OrderPaidError
	}

	// 标记订单
	if err = orderQuery.MarkOrderPaid(req.OrderId); err != nil {
		klog.Errorf("order of orderId %s mark paid err: %v", req.OrderId, err)
		return nil, kitex_err.MysqlError
	}

	// 返回
	return &order.MarkOrderPaidResp{}, nil
}
