package order

import (
	"byte_go/backend/app/front/infra/rpc"
	rpcOrder "byte_go/backend/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	order "byte_go/backend/app/front/hertz_gen/frontend/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type MarkOrderPaidService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMarkOrderPaidService(Context context.Context, RequestContext *app.RequestContext) *MarkOrderPaidService {
	return &MarkOrderPaidService{RequestContext: RequestContext, Context: Context}
}

func (h *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {

	_, err = rpc.OrderClient.MarkOrderPaid(h.Context, &rpcOrder.MarkOrderPaidReq{
		OrderId: req.OrderId,
		UserId:  req.UserId,
	})

	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return
}
