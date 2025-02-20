package service

import (
	"byte_go/backend/app/cart/biz/dal/mysql"
	"byte_go/backend/app/cart/biz/model"
	cart "byte_go/backend/rpc_gen/kitex_gen/cart"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {

	cartQuery := model.NewCartQuery(s.ctx, mysql.DB)
	err = cartQuery.AddItem(&model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	})

	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}

	return
}
