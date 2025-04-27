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

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// 校验参数
	if req == nil {
		return nil, kitex_err.RequestParamError
	}
	if len(req.Query) > 100 {
		req.Query = req.Query[:100]
	}

	// 从数据库中查询商品
	productQuery := repository.NewProductRepository(s.ctx, mysql.DB, redis.RedisClient)
	products, err := productQuery.SearchProducts(req.Query)
	if err != nil {
		klog.Errorf("search products failed: %v", err.Error())
		return nil, kitex_err.MysqlError
	}

	// 封装商品
	resp = &product.SearchProductsResp{}
	resp.Products = make([]*product.Product, len(products))
	for idx, p := range products {
		resp.Products[idx] = model.ProductModel2Gen(p)
	}

	// 返回
	return resp, nil
}
