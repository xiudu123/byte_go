package service

import (
	"byte_go/backend/app/cart/biz/dal/mysql"
	"byte_go/backend/app/cart/biz/model"
	cart "byte_go/backend/rpc_gen/kitex_gen/cart"
	"byte_go/kitex_err"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {

	// 获取购物车
	cartQuery := model.NewCartQuery(s.ctx, mysql.DB)
	cartItems, err := cartQuery.GetCartByUserId(req.UserId)

	// 处理错误
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &cart.GetCartResp{
				Cart: &cart.Cart{
					UserId: req.UserId,
				},
			}, nil
		}
		return nil, kitex_err.SystemError
	}

	// 获取商品
	resp = &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
		},
	}
	resp.Cart.Items = make([]*cart.CartItem, len(cartItems))
	for idx, item := range cartItems {
		resp.Cart.Items[idx] = &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}

	// 返回
	return resp, nil
}
