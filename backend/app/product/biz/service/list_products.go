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

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// 校验参数
	if req == nil {
		return nil, kitex_err.RequestParamError
	}
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	// 从数据库中查询商品
	productQuery := repository.NewProductRepository(s.ctx, mysql.DB, redis.RedisClient)
	productCount, products, err := productQuery.ListProductsByCategory(req.Page, int32(req.PageSize), req.CategoryName)

	if err != nil {
		klog.Errorf("list products failed: %v", err.Error())
		return nil, kitex_err.MysqlError
	}

	// 封装商品
	resp = &product.ListProductsResp{}
	resp.Products = make([]*product.Product, len(products))
	resp.Total = productCount
	for idx, p := range products {
		resp.Products[idx] = model.ProductModel2Gen(p)
	}

	// 返回
	return resp, nil
}
