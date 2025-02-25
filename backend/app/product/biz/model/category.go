package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/2/14 22:25
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`

	Products []Product `gorm:"many2many:product_category;"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}

func (c CategoryQuery) CreateCategory(category Category) (err error) {
	err = c.db.WithContext(c.ctx).Create(&category).Error
	return
}

func (c CategoryQuery) GetCategoriesByNames(categoryNames []string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Where("name IN ?", categoryNames).Find(&categories).Error
	return
}

func (c CategoryQuery) ExistCateGoryByName(categoryName string) (bool, error) {
	err := c.db.WithContext(c.ctx).Where("name =?", categoryName).First(&Category{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {

		return false, err
	}
	return true, nil
}