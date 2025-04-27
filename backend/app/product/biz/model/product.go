package model

import (
	"byte_go/backend/rpc_gen/kitex_gen/product"
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

// ProductModel2Gen 将 Product 模型转换为 gen.Product 类型
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
