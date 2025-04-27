package service

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/app/product/biz/dal/redis"
	"byte_go/backend/app/product/biz/dal/repository"
	"byte_go/backend/app/product/biz/model"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListProductByIdsService struct {
	ctx context.Context
} // NewListProductByIdsService new ListProductByIdsService
func NewListProductByIdsService(ctx context.Context) *ListProductByIdsService {
	return &ListProductByIdsService{ctx: ctx}
}

// Run create note info
func (s *ListProductByIdsService) Run(req *product.ListProductByIdsReq) (resp *product.ListProductByIdsResp, err error) {

	// 校验参数
	if req == nil {
		return nil, kitex_err.RequestParamError
	}
	if len(req.ProductIds) == 0 {
		return nil, kitex_err.ProductEmptyError
	}

	// 从数据库中查询商品
	productQuery := repository.NewProductRepository(s.ctx, mysql.DB, redis.RedisClient)
	productIds := make([]uint, len(req.ProductIds))
	for idx, id := range req.ProductIds {
		productIds[idx] = uint(id)
	}
	productResult, err := productQuery.ListProductsByIds(productIds)
	if err != nil {
		klog.Errorf("list products by ids failed: %v", err.Error())
		return nil, kitex_err.MysqlError
	}

	// 封装商品
	products := make([]*product.Product, len(productResult))
	for idx, p := range productResult {
		products[idx] = model.ProductModel2Gen(p)
	}

	// 返回
	return &product.ListProductByIdsResp{
		Products: products,
	}, nil
}
