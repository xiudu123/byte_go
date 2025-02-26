package cart

import (
	"byte_go/backend/app/front/biz/utils"
	cart "byte_go/backend/app/front/hertz_gen/frontend/cart"
	"byte_go/backend/app/front/infra/rpc"
	rpcCart "byte_go/backend/rpc_gen/kitex_gen/cart"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type AddItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddItemService(Context context.Context, RequestContext *app.RequestContext) *AddItemService {
	return &AddItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {

	_, err = rpc.CartClient.AddItem(h.Context, &rpcCart.AddItemReq{UserId: req.UserId, Item: utils.CartItemHertz2Gen(req.Item)})

	if err != nil {
		hlog.CtxErrorf(h.Context, "add item [%v] failed, err: %v", req.Item, err.Error())
		return nil, err
	}

	return
}
