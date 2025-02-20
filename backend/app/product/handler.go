package main

import (
	"byte_go/backend/app/product/biz/service"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"context"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// ListProductByIds implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProductByIds(ctx context.Context, req *product.ListProductByIdsReq) (resp *product.ListProductByIdsResp, err error) {
	resp, err = service.NewListProductByIdsService(ctx).Run(req)

	return resp, err
}
