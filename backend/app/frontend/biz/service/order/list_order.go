package order

import (
	"byte_go/backend/app/front/biz/utils"
	order "byte_go/backend/app/front/hertz_gen/frontend/order"
	"byte_go/backend/app/front/infra/rpc"
	rpcOrder "byte_go/backend/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ListOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListOrderService(Context context.Context, RequestContext *app.RequestContext) *ListOrderService {
	return &ListOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	orderList, err := rpc.OrderClient.ListOrder(h.Context, &rpcOrder.ListOrderReq{
		UserId: req.UserId,
	})

	if err != nil {
		hlog.CtxErrorf(h.Context, "list order failed, err: %v", err.Error())
		return nil, err
	}

	return &order.ListOrderResp{
		Orders: utils.OrderListGen2Hertz(orderList.Orders),
	}, nil
}
