package cart

import (
	cart2 "byte_go/backend/app/frontend/biz/service/cart"
	"context"

	"byte_go/backend/app/frontend/biz/utils"
	cart "byte_go/backend/app/frontend/hertz_gen/frontend/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AddItem .
// @router /cart/add [POST]
func AddItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &cart.AddItemResp{}
	resp, err = cart2.NewAddItemService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetCart .
// @router /cart/get [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.GetCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &cart.GetCartResp{}
	resp, err = cart2.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// EmptyCart .
// @router /cart/empty [POST]
func EmptyCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.EmptyCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &cart.EmptyCartResp{}
	resp, err = cart2.NewEmptyCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
