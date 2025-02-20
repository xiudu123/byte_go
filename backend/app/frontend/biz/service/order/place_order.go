package order

import (
	"context"

	order "byte_go/backend/app/front/hertz_gen/frontend/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type PlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *PlaceOrderService {
	return &PlaceOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {

	return
}
