package product

import (
	"byte_go/backend/app/front/biz/utils"
	"byte_go/backend/app/front/infra/rpc"
	rpcProduct "byte_go/backend/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	product "byte_go/backend/app/front/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListProductsService(Context context.Context, RequestContext *app.RequestContext) *ListProductsService {
	return &ListProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {

	// 请求rpc
	productResult, err := rpc.ProductClient.ListProducts(h.Context, &rpcProduct.ListProductsReq{
		Page:         req.Page,
		PageSize:     req.PageSize,
		CategoryName: req.CategoryName,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "list products failed, err: %v", err.Error())
		return nil, err
	}

	// 封装数据
	resp = &product.ListProductsResp{}
	resp.Products = make([]*product.Product, len(productResult.Products))
	for idx, p := range productResult.Products {
		resp.Products[idx] = utils.ProductGen2Hertz(p)
	}

	// 返回
	return resp, nil
}
