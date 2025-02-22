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

	// 参数校验
	if req == nil || req.UserId == 0 || req.Item.ProductId == 0 || req.Item.Quantity <= 0 {
		return nil, kitex_err.RequestParamError
	}

	// 数据库操作
	cartQuery := model.NewCartQuery(s.ctx, mysql.DB)
	err = cartQuery.AddItem(&model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	})
	if err != nil {
		klog.Errorf("user of userId %d add item of productId %d err: %v", req.UserId, req.Item.ProductId, err)
		return nil, kitex_err.MysqlError
	}

	// 返回
	return &cart.AddItemResp{}, nil
}
