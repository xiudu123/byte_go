package service

import (
	"byte_go/backend/app/cart/biz/dal/mysql"
	"byte_go/backend/app/cart/biz/model"
	cart "byte_go/backend/rpc_gen/kitex_gen/cart"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {

	// 参数校验
	if req == nil || req.UserId == 0 {
		return nil, kitex_err.RequestParamError
	}

	// 数据库操作
	cartQuery := model.NewCartQuery(s.ctx, mysql.DB)
	err = cartQuery.EmptyCartByUserId(req.UserId)
	if err != nil {
		klog.Errorf("user of userId %d empty cart err: %v", req.UserId, err)
		return nil, kitex_err.MysqlError
	}
	// 返回
	return &cart.EmptyCartResp{}, nil
}
