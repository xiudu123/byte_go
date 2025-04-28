package service

import (
	"byte_go/backend/app/product/biz/dal/repository"
	"byte_go/backend/app/product/biz/model"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CreateCategoryService struct {
	ctx context.Context
} // NewCreateCategoryService new CreateCategoryService
func NewCreateCategoryService(ctx context.Context) *CreateCategoryService {
	return &CreateCategoryService{ctx: ctx}
}

// Run create note info
func (s *CreateCategoryService) Run(req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {

	// 校验参数
	if req == nil || req.Name == "" {
		return nil, kitex_err.RequestParamError
	}

	categoryQuery := repository.NewCategoryRepository(s.ctx)

	// 判断分类是否存在
	categoryExist, err := categoryQuery.ExistCategoryByName(req.Name)
	if err != nil {
		klog.Errorf("check category [%s] exist failed: %v", req.Name, err.Error())
		return nil, kitex_err.MysqlError
	}
	if categoryExist {
		return nil, kitex_err.CategoryExist
	}

	// 创建分类
	category := model.Category{
		Name: req.Name,
	}
	err = categoryQuery.CreateCategory(&category)
	if err != nil {
		klog.Errorf("create category [%s] failed: %v", req.Name, err.Error())
		return nil, kitex_err.MysqlError
	}
	// 返回
	return &product.CreateCategoryResp{
		CategoryId: uint32(category.ID),
	}, nil
}
