package service

import (
	"byte_go/backend/app/order/biz/dal/mysql"
	"byte_go/backend/app/order/biz/model"
	order "byte_go/backend/rpc_gen/kitex_gen/order"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {

	orderQuery := model.NewOrderQuery(s.ctx, mysql.DB)
	if err = orderQuery.MarkOrderPaid(req.OrderId); err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}
	return
}
