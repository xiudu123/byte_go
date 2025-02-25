package product

import (
	"byte_go/backend/app/front/infra/rpc"
	rpcProduct "byte_go/backend/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	product "byte_go/backend/app/front/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateCategoryService(Context context.Context, RequestContext *app.RequestContext) *CreateCategoryService {
	return &CreateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateCategoryService) Run(req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {

	// 请求rpc
	categoryResult, err := rpc.ProductClient.CreateCategory(h.Context, &rpcProduct.CreateCategoryReq{
		Name: req.Name,
	})
	if err != nil {
		hlog.Error(err)
		return
	}
	// 封装返回
	return &product.CreateCategoryResp{
		CategoryId: categoryResult.CategoryId,
	}, nil
}
