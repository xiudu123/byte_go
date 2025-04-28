package repository

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/app/product/biz/dal/mysql/dao"
	"byte_go/backend/app/product/biz/model"
	"context"
)

/**
 * @author: 锈渎
 * @date: 2025/4/27 21:02
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type CategoryRepository struct {
	ctx context.Context
	dao *dao.CategoryDAO
}

func NewCategoryRepository(ctx context.Context) *CategoryRepository {
	return &CategoryRepository{
		ctx: ctx,
		dao: dao.NewCategoryDAO(mysql.DB),
	}
}

// CreateCategory 创建分类
func (c *CategoryRepository) CreateCategory(category *model.Category) (err error) {
	err = c.dao.CreateCategory(c.ctx, category)
	return
}

// GetCategoriesByNames 获取分类
func (c *CategoryRepository) GetCategoriesByNames(categoryNames []string) (categories []model.Category, err error) {

	categories, err = c.dao.GetCategoriesByNames(c.ctx, categoryNames)
	return
}

// ExistCategoryByName 判断分类是否存在
func (c *CategoryRepository) ExistCategoryByName(categoryName string) (bool, error) {
	exist, err := c.dao.ExistCategoryByName(c.ctx, categoryName)
	return exist, err
}
