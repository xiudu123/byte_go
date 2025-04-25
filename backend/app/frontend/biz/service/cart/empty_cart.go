package cart

import (
	cart "byte_go/backend/app/frontend/hertz_gen/frontend/cart"
	"byte_go/backend/app/frontend/infra/rpc"
	rpcCart "byte_go/backend/rpc_gen/kitex_gen/cart"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	_, err = rpc.CartClient.EmptyCart(h.Context, &rpcCart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		hlog.CtxErrorf(h.Context, "empty cart failed, err: %v", err.Error())
		return nil, err
	}
	return
}
