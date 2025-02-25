package service

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/app/product/biz/model"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {

	// 校验参数
	if req == nil || req.ProductId <= 0 {
		return nil, kitex_err.RequestParamError
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	// 检查商品是否存在
	productExist, err := productQuery.ExistProductById(uint(req.ProductId))
	if err != nil {
		klog.Errorf("check product [%d] exist failed: %v", req.ProductId, err.Error())
		return nil, kitex_err.MysqlError
	}
	if !productExist {
		return nil, kitex_err.ProductNotExist
	}

	// 删除商品
	err = productQuery.DeleteProduct(uint(req.ProductId))
	if err != nil {
		klog.Errorf("delete product [%d] failed: %v", req.ProductId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 返回
	return &product.DeleteProductResp{}, nil
}
