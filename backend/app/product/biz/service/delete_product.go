package service

import (
	"byte_go/backend/app/product/biz/dal/repository"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
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
	// 定义查询对象
	productQuery := repository.NewProductRepository(s.ctx)

	// 检查商品是否存在
	productOld, err := productQuery.GetProductById(uint(req.ProductId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kitex_err.ProductNotExist
		}
		klog.Errorf("check product [%d] exist failed: %v", req.ProductId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 清空商品分类
	if err = productQuery.ClearCategory(productOld); err != nil {
		klog.Errorf("clear product [%d] category failed: %v", req.ProductId, err.Error())
		return nil, kitex_err.MysqlError
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
