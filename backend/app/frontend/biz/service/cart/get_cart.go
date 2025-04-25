package cart

import (
	"byte_go/backend/app/frontend/biz/utils"
	"byte_go/backend/app/frontend/infra/rpc"
	rpcCart "byte_go/backend/rpc_gen/kitex_gen/cart"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	cart "byte_go/backend/app/frontend/hertz_gen/frontend/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {

	cartResult, err := rpc.CartClient.GetCart(h.Context, &rpcCart.GetCartReq{UserId: req.UserId})

	if err != nil {
		hlog.CtxErrorf(h.Context, "get cart failed, err: %v", err.Error())
		return nil, err
	}

	return &cart.GetCartResp{
		Cart: utils.CartGen2Hertz(cartResult.Cart),
	}, nil
}
