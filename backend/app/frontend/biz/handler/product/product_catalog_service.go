package product

import (
	product2 "byte_go/backend/app/front/biz/service/product"
	"context"

	"byte_go/backend/app/front/biz/utils"
	product "byte_go/backend/app/front/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListProducts .
// @router /products/list [GET]
func ListProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ListProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.ListProductsResp{}
	resp, err = product2.NewListProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetProduct .
// @router /products/get/:product_id [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.GetProductResp{}
	resp, err = product2.NewGetProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// SearchProducts .
// @router /products/search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &product.SearchProductsResp{}
	resp, err = product2.NewSearchProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateProduct .
// @router /products/create [POST]
func CreateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCreateProductService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteProduct .
// @router /products/delete/:product_id [POST]
func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.DeleteProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewDeleteProductService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateProduct .
// @router /products/update/:product_id [POST]
func UpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUpdateProductService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateCategory .
// @router /products/create_category [POST]
func CreateCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCreateCategoryService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
