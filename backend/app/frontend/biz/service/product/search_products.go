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

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {

	// 请求rpc
	productResult, err := rpc.ProductClient.SearchProducts(h.Context, &rpcProduct.SearchProductsReq{
		Query: req.Query,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "search products [%s] failed, err: %v", req.Query, err.Error())
		return nil, err
	}

	resp = &product.SearchProductsResp{}
	// 封装数据
	resp.Products = make([]*product.Product, len(productResult.Products))
	for idx, p := range productResult.Products {
		resp.Products[idx] = utils.ProductGen2Hertz(p)
	}

	// 返回
	return resp, nil
}
