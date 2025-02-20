package service

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/app/product/biz/model"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// 校验参数
	if req.ProductId <= 0 {
		return nil, kitex_err.ProductNotExist
	}

	// 获取商品
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.GetProductById(uint(req.ProductId))
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}
	categoryNames := make([]string, len(p.Categories))
	for idx, c := range p.Categories {
		categoryNames[idx] = c.Name
	}

	// 返回
	return &product.GetProductResp{
		Product: &product.Product{
			ProductId:   uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  categoryNames,
		},
	}, nil
}
