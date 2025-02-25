package product

import (
	"byte_go/backend/app/front/infra/rpc"
	rpcProduct "byte_go/backend/rpc_gen/kitex_gen/product"
	"context"

	product "byte_go/backend/app/front/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	productResult, err := rpc.ProductClient.UpdateProduct(h.Context, &rpcProduct.UpdateProductReq{
		ProductId:   req.ProductId,
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	})
	if err != nil {
		return nil, err
	}

	return &product.UpdateProductResp{
		ProductId: productResult.ProductId,
	}, nil
}
