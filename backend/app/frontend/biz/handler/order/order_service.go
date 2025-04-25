package order

import (
	order2 "byte_go/backend/app/frontend/biz/service/order"
	"context"

	"byte_go/backend/app/frontend/biz/utils"
	order "byte_go/backend/app/frontend/hertz_gen/frontend/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PlaceOrder .
// @router /order/place [GET]
func PlaceOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.PlaceOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.PlaceOrderResp{}
	resp, err = order2.NewPlaceOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListOrder .
// @router /order/list [GET]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.ListOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.ListOrderResp{}
	resp, err = order2.NewListOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MarkOrderPaid .
// @router /order/mark_paid [POST]
func MarkOrderPaid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.MarkOrderPaidReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := order2.NewMarkOrderPaidService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
