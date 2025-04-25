package payment

import (
	payment2 "byte_go/backend/app/frontend/biz/service/payment"
	"context"

	"byte_go/backend/app/frontend/biz/utils"
	payment "byte_go/backend/app/frontend/hertz_gen/frontend/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Charge .
// @router /payment/charge [POST]
func Charge(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.ChargeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.ChargeResp{}
	resp, err = payment2.NewChargeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
