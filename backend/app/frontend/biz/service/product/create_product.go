package product

import (
	"byte_go/backend/app/front/infra/rpc"
	rpcProduct "byte_go/backend/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	product "byte_go/backend/app/front/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProductService(Context context.Context, RequestContext *app.RequestContext) *CreateProductService {
	return &CreateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// 请求rpc
	productResult, err := rpc.ProductClient.CreateProduct(h.Context, &rpcProduct.CreateProductReq{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	})
	if err != nil {
		hlog.Error(err)
		return
	}
	// 封装返回
	return &product.CreateProductResp{
		ProductId: productResult.ProductId,
	}, nil
}
