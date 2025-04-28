package dao

import (
	"byte_go/backend/app/product/biz/model"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/4/27 20:52
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type CategoryDAO struct {
	db *gorm.DB
}

func NewCategoryDAO(db *gorm.DB) *CategoryDAO {
	return &CategoryDAO{
		db: db,
	}
}

// CreateCategory 创建分类
func (dao *CategoryDAO) CreateCategory(ctx context.Context, category *model.Category) (err error) {
	if err := dao.db.WithContext(ctx).Create(category).Error; err != nil {
		klog.Errorf("mysql: create category failed: %v", err.Error())
		return err
	}
	return nil
}

// GetCategoriesByNames 获取分类
func (dao *CategoryDAO) GetCategoriesByNames(ctx context.Context, categoryNames []string) (categories []model.Category, err error) {
	if err := dao.db.WithContext(ctx).Where("name IN ?", categoryNames).Find(&categories).Error; err != nil {
		klog.Errorf("mysql: get categories by names failed: %v", err.Error())
		return nil, err
	}
	return categories, nil
}

// ExistCategoryByName 判断分类是否存在
func (dao *CategoryDAO) ExistCategoryByName(ctx context.Context, categoryName string) (bool, error) {
	var count int64
	if err := dao.db.WithContext(ctx).Model(&model.Category{}).Where("name =?", categoryName).Count(&count).Error; err != nil {
		klog.Errorf("mysql: check category [%s] exist failed: %v", categoryName, err.Error())
		return false, err
	}
	return count > 0, nil
}
