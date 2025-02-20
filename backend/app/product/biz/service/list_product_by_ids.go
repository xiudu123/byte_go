package service

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/app/product/biz/model"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
)

type ListProductByIdsService struct {
	ctx context.Context
} // NewListProductByIdsService new ListProductByIdsService
func NewListProductByIdsService(ctx context.Context) *ListProductByIdsService {
	return &ListProductByIdsService{ctx: ctx}
}

// Run create note info
func (s *ListProductByIdsService) Run(req *product.ListProductByIdsReq) (resp *product.ListProductByIdsResp, err error) {

	if len(req.ProductIds) == 0 {
		return nil, kitex_err.ProductEmptyError
	}

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	productIds := make([]uint, len(req.ProductIds))
	for idx, id := range req.ProductIds {
		productIds[idx] = uint(id)
	}
	productResult, err := productQuery.ListProductsByIds(productIds)

	products := make([]*product.Product, len(productResult))
	for idx, p := range productResult {
		products[idx] = model.ProductModel2Gen(&p)
	}

	return &product.ListProductByIdsResp{
		Products: products,
	}, nil
}
