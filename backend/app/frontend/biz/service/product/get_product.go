package product

import (
	product "byte_go/backend/app/front/hertz_gen/frontend/product"
	"byte_go/backend/app/front/infra/rpc"
	rpcProduct "byte_go/backend/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {

	// 请求rpc
	productResult, err := rpc.ProductClient.GetProduct(h.Context, &rpcProduct.GetProductReq{ProductId: req.ProductId})

	// 封装返回
	if err != nil {
		hlog.CtxErrorf(h.Context, "get product [%d] failed, err: %v", req.ProductId, err.Error())
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			ProductId:   productResult.Product.ProductId,
			Name:        productResult.Product.Name,
			Description: productResult.Product.Description,
			Picture:     productResult.Product.Picture,
			Price:       productResult.Product.Price,
			Categories:  productResult.Product.Categories,
		},
	}, nil
}
