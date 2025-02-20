package service

import (
	"byte_go/backend/app/product/biz/dal/mysql"
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

	// 从数据库中查询商品
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}

	resp = &product.SearchProductsResp{}
	// 封装商品
	resp.Products = make([]*product.Product, len(products))
	for idx, p := range products {
		resp.Products[idx] = model.ProductModel2Gen(&p)
	}

	// 返回
	return resp, nil
}
