package order

import (
	"byte_go/backend/app/front/infra/rpc"
	rpcCart "byte_go/backend/rpc_gen/kitex_gen/cart"
	rpcOrder "byte_go/backend/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

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

	orderItems := make([]*rpcOrder.OrderItem, len(req.OrderItems))
	for i, item := range req.OrderItems {
		orderItems[i] = &rpcOrder.OrderItem{
			Item: &rpcCart.CartItem{
				ProductId: item.Item.ProductId,
				Quantity:  item.Item.Quantity,
			},
			Cost: item.Cost,
		}
	}

	orderResult, err := rpc.OrderClient.PlaceOrder(h.Context, &rpcOrder.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Email:        req.Email,
		Address: &rpcOrder.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			ZipCode:       req.Address.ZipCode,
			Country:       req.Address.Country,
		},
		OrderItems: orderItems,
		Nickname:   req.Nickname,
	})

	if err != nil {
		hlog.CtxErrorf(h.Context, "place order failed, err: %v", err.Error())
		return nil, err
	}

	return &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: orderResult.Order.OrderId,
		},
	}, nil
}
