package service

import (
	"byte_go/backend/app/product/biz/dal/mysql"
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

	// 从数据库中查询商品
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	productCount, products, err := productQuery.ListProductsByCategory(req.Page, req.PageSize, req.CategoryName)

	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}

	// 封装商品
	resp = &product.ListProductsResp{}
	resp.Products = make([]*product.Product, len(products))
	resp.Total = productCount
	for idx, p := range products {
		resp.Products[idx] = model.ProductModel2Gen(&p)
	}

	// 返回
	return resp, nil
}
