package service

import (
	"byte_go/backend/app/product/biz/dal/repository"
	"byte_go/backend/app/product/biz/model"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {

	// 校验参数
	if req == nil || req.Name == "" || req.Description == "" || req.Price <= 0 || len(req.Picture) == 0 || len(req.Categories) == 0 {
		return nil, kitex_err.RequestParamError
	}

	productQuery := repository.NewProductRepository(s.ctx)
	categoryQuery := repository.NewCategoryRepository(s.ctx)
	// 检查商品是否存在
	productOld, err := productQuery.GetProductById(uint(req.ProductId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kitex_err.ProductNotExist
		}
		klog.Errorf("check product [%d] exist failed: %v", req.ProductId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 清空商品分类
	if err = productQuery.ClearCategory(productOld); err != nil {
		klog.Errorf("clear product [%d] category failed: %v", req.ProductId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 获取分类
	categories, err := categoryQuery.GetCategoriesByNames(req.Categories)
	if err != nil {
		klog.Errorf("get categories [%v] failed: %v", req.Categories, err.Error())
		return nil, kitex_err.MysqlError
	}
	if len(categories) != len(req.Categories) {
		return nil, kitex_err.CategoryNotExist
	}

	// 更新商品
	productReq := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  categories,
	}
	productReq.ID = uint(req.ProductId)
	err = productQuery.UpdateProduct(productReq)
	if err != nil {
		klog.Errorf("update product [%d] failed: %v", req.ProductId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 返回
	return &product.UpdateProductResp{
		ProductId: req.ProductId,
	}, nil
}
