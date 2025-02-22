package service

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/app/product/biz/model"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kitex_err.ProductNotExist
	}
	if err != nil {
		klog.Errorf("get product by id '%d' failed: %v", req.ProductId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 处理商品分类
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
