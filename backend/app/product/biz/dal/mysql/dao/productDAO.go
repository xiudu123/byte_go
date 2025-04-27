package dao

import (
	"byte_go/backend/app/product/biz/model"
	"context"
	"errors"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"gorm.io/gorm"
)

/**
 * @author: 锈渎
 * @date: 2025/4/27 15:37
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

var ErrRecordNotFound = errors.New("record not found")

type ProductDAO struct {
	db *gorm.DB
}

func NewProductDAO(db *gorm.DB) *ProductDAO {
	return &ProductDAO{
		db: db,
	}
}

// GetByID 根据ID查询商品
func (dao *ProductDAO) GetByID(ctx context.Context, productId uint) (*model.Product, error) {
	var product *model.Product

	err := dao.db.
		WithContext(ctx).
		Preload("Categories").
		Where("id =?", productId).
		First(&product).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Infof("mysql: product by id [%d] not found", productId)
			return &model.Product{}, ErrRecordNotFound
		}
		log.Errorf("mysql: product by id [%d] failed: %v", productId, err.Error())
		return nil, err
	}

	return product, nil
}

// ListByIDs 根据ID列表查询商品
func (dao *ProductDAO) ListByIDs(ctx context.Context, productIds []uint) ([]*model.Product, error) {
	var products []*model.Product

	err := dao.db.
		WithContext(ctx).
		Preload("Categories").
		Where("id IN ?", productIds).
		Find(&products).Error

	if err != nil {
		log.Errorf("mysql: product by ids [%v] failed: %v", productIds, err.Error())
		return nil, err
	}

	return products, nil
}

// Create 创建商品
func (dao *ProductDAO) Create(ctx context.Context, product *model.Product) error {
	if err := dao.db.WithContext(ctx).Create(&product).Error; err != nil {
		log.Errorf("mysql: create product failed: %v", err.Error())
		return err
	}
	return nil
}

// Update 更新商品
func (dao *ProductDAO) Update(ctx context.Context, product *model.Product) error {
	result := dao.db.WithContext(ctx).Model(&product).Updates(&product)

	if result.Error != nil {
		log.Errorf("mysql: update product failed: %v", result.Error.Error())
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("mysql: update product failed: product by id [%d] not exist", product.ID)
		return ErrRecordNotFound
	}

	return nil
}

// Delete 删除商品
func (dao *ProductDAO) Delete(ctx context.Context, productId uint) error {
	result := dao.db.WithContext(ctx).Delete(&model.Product{}, productId)

	if result.Error != nil {
		log.Errorf("mysql: delete product failed: %v", result.Error.Error())
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("mysql: delete product failed: product by id [%d] not exist", productId)
		return ErrRecordNotFound
	}

	return nil
}

// ClearCategories 清除商品与分类的关联关系
// 注意：这里使用的是GORM的Association来处理多对多关系
func (dao *ProductDAO) ClearCategories(ctx context.Context, productId uint) error {
	product := &model.Product{
		Model: gorm.Model{
			ID: productId,
		},
	}

	if err := dao.db.WithContext(ctx).Model(&product).Association("Categories").Clear(); err != nil {
		log.Errorf("mysql: clear product categories failed: %v", err.Error())
		return err
	}
	return nil
}

// ListByCategory 根据分类查询商品
func (dao *ProductDAO) ListByCategory(ctx context.Context, categoryName string, page, pageSize int) ([]*model.Product, int64, error) {
	var (
		products []*model.Product
		count    int64
	)

	// 构建基础查询
	query := dao.db.
		WithContext(ctx).
		Model(&model.Product{}).
		Joins("JOIN product_category ON product_category.product_id = product.id").
		Joins("JOIN category ON category.id = product_category.category_id").
		Where("category.name = ?", categoryName)
	//Distinct("product.id") // 联表查询，去重

	// 计算总数
	if err := query.Count(&count).Error; err != nil {
		log.Errorf("mysql: list product by category failed: %v", err.Error())
		return nil, 0, err
	}

	// 计算偏移量
	offset := (page - 1) * pageSize
	if int64(offset) > count {
		return nil, count, nil
	}

	// 分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		log.Errorf("mysql: list product by category failed: %v", err.Error())
		return nil, 0, err
	}
	return products, count, nil
}

// Search 搜索商品
func (dao *ProductDAO) Search(ctx context.Context, query string) ([]*model.Product, error) {
	var products []*model.Product
	searchPattern := "%" + query + "%"

	if err := dao.db.WithContext(ctx).
		Where("name LIKE ?", searchPattern).
		Or("description LIKE ?", searchPattern).
		Find(&products).Error; err != nil {
		log.Errorf("mysql: search product failed: %v", err.Error())
		return nil, err
	}
	return products, nil
}
