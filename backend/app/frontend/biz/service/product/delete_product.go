package product

import (
	"byte_go/backend/app/front/infra/rpc"
	rpcProduct "byte_go/backend/rpc_gen/kitex_gen/product"
	"context"

	product "byte_go/backend/app/front/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {

	_, err = rpc.ProductClient.DeleteProduct(h.Context, &rpcProduct.DeleteProductReq{
		ProductId: req.ProductId,
	})
	if err != nil {
		return nil, err
	}

	return &product.DeleteProductResp{}, nil
}
