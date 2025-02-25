package model

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/rpc_gen/kitex_gen/product"
	"context"
	"errors"
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/2/14 22:23
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:varchar(255);not null"`
	Picture     string  `gorm:"type:varchar(255);not null"`
	Price       float32 `gorm:"type:decimal(10,2);not null"`

	Categories []Category `gorm:"many2many:product_category;"`
}

func (p Product) TableName() string {
	return "product"
}

func ProductModel2Gen(modelProduct *Product) (genProduct *product.Product) {
	categoryNames := make([]string, len(modelProduct.Categories))
	for i, c := range modelProduct.Categories {
		categoryNames[i] = c.Name
	}
	genProduct = &product.Product{
		ProductId:   uint32(modelProduct.ID),
		Name:        modelProduct.Name,
		Description: modelProduct.Description,
		Picture:     modelProduct.Picture,
		Price:       modelProduct.Price,
		Categories:  categoryNames,
	}
	return
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

func (p ProductQuery) GetProductById(productId uint) (product Product, err error) {

	err = p.db.WithContext(p.ctx).Preload("Categories").Where("id = ?", productId).First(&product).Error
	return
}

func (p ProductQuery) ListProductsByCategory(page int32, pageSize int64, categoryName string) (count int64, products []Product, err error) {

	// 构建基础查询
	query := p.db.WithContext(p.ctx).
		Model(&Product{}).
		Joins("JOIN product_category ON product_category.product_id = product.id").
		Joins("JOIN category ON category.id = product_category.category_id").
		Where("category.name = ?", categoryName)
	//Distinct("product.id") // 联表查询，去重

	// 计算总数
	err = query.Count(&count).Error
	if err != nil {
		return
	}

	// 计算偏移量
	offset := int64(page-1) * pageSize
	if offset > count {
		return count, nil, nil
	}

	// 分页查询
	err = query.
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&products).Error

	return
}

func (p ProductQuery) SearchProducts(query string) (products []Product, err error) {
	err = p.db.WithContext(p.ctx).
		Where("name LIKE ?", "%"+query+"%").
		Or("description LIKE ?", "%"+query+"%").
		Find(&products).Error
	return
}

func (p ProductQuery) ListProductsByIds(productIds []uint) (products []Product, err error) {
	err = p.db.WithContext(p.ctx).
		Where("id IN ?", productIds).
		Find(&products).Error
	return
}

func (p ProductQuery) CreateProduct(product Product) (err error) {
	// 开启事务
	tx := mysql.DB.Begin()
	err = p.db.WithContext(p.ctx).Create(&product).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return
}

func (p ProductQuery) DeleteProduct(productId uint) (err error) {
	err = p.db.WithContext(p.ctx).Delete(&Product{}, productId).Error
	return

}

func (p ProductQuery) UpdateProduct(product Product) (err error) {
	err = p.db.WithContext(p.ctx).Save(&product).Error
	return
}

func (p ProductQuery) ExistProductById(productId uint) (bool, error) {
	err := p.db.WithContext(p.ctx).Where("id =?", productId).First(&Product{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
