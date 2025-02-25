package service

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/app/product/biz/model"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {

	// 校验参数
	if req == nil || req.Name == "" || req.Description == "" || req.Price <= 0 || len(req.Picture) == 0 || len(req.Categories) == 0 {
		return nil, kitex_err.RequestParamError
	}

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	// 处理商品分类
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	categories, err := categoryQuery.GetCategoriesByNames(req.Categories)
	if err != nil {
		klog.Errorf("get categories failed: %v", err.Error())
		return nil, kitex_err.MysqlError
	}
	if len(categories) != len(req.Categories) {
		return nil, kitex_err.CategoryNotExist
	}
	// 插入商品
	productReq := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  categories,
	}
	err = productQuery.CreateProduct(productReq)
	if err != nil {
		klog.Errorf("create product failed: %v", err.Error())
		return nil, kitex_err.MysqlError
	}

	// 返回
	return &product.CreateProductResp{
		ProductId: uint32(productReq.ID),
	}, nil
}
