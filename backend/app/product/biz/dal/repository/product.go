package repository

import (
	"byte_go/backend/app/product/biz/dal/mysql/dao"
	"byte_go/backend/app/product/biz/dal/redis/cache"
	"byte_go/backend/app/product/biz/model"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/4/27 14:53
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type ProductRepository struct {
	ctx          context.Context
	productGroup *singleflight.Group // 合并相同请求，防止缓存击穿
	cache        *cache.ProductCache // 缓存实例
	dao          *dao.ProductDAO     // 数据库操作实例
}

// NewProductRepository 创建新的 ProductRepository 实例
// ctx: 上下文
// db: GORM 数据库连接
// redisClient: Redis 客户端
func NewProductRepository(ctx context.Context, db *gorm.DB, redisClient *redis.Client) *ProductRepository {
	return &ProductRepository{
		ctx:          ctx,
		productGroup: &singleflight.Group{},
		cache:        cache.NewProductCache(redisClient), // 初始化缓存实例
		dao:          dao.NewProductDAO(db),              // 初始化数据库操作实例
	}
}

// GetProductById 根据ID查询商品
func (p *ProductRepository) GetProductById(productId uint) (*model.Product, error) {

	// 从缓存中获取数据
	if cacheProduct, err := p.cache.Get(p.ctx, productId); err == nil {
		return cacheProduct, nil
	}

	// 使用SingleFlight模式，确保同一时间只有一个请求查询数据库, 防止缓存击穿
	result, err, _ := p.productGroup.Do("get product from mysql "+p.cache.BuildProductCacheKey(productId), func() (interface{}, error) {
		// 从数据库查询数据
		dbProduct, err := p.dao.GetByID(p.ctx, productId)
		if err != nil {
			// 数据库不存在数据
			if errors.Is(err, dao.ErrRecordNotFound) {
				go func() {
					_ = p.cache.SetNotFound(context.Background(), productId) // 缓存空值，防止缓存穿透
				}()
				return &model.Product{}, nil
			}

			return nil, err
		}

		// 异步更新缓存
		go func() { // 使用协程避免阻塞主流程
			_ = p.cache.Set(context.Background(), dbProduct)
		}()
		return dbProduct, nil
	})

	if err != nil {
		return nil, err
	}
	return result.(*model.Product), nil
}

// ListProductsByCategory 根据分类查询商品
func (p *ProductRepository) ListProductsByCategory(page int32, pageSize int32, categoryName string) (count int64, products []*model.Product, err error) {
	// TODO: 用es实现
	products, count, err = p.dao.ListByCategory(p.ctx, categoryName, int(page), int(pageSize))
	return
}

// SearchProducts 搜索商品
func (p *ProductRepository) SearchProducts(query string) (products []*model.Product, err error) {
	// TODO: 用es实现
	products, err = p.dao.Search(p.ctx, query)
	return
}

// ListProductsByIds 根据ID列表查询商品
func (p *ProductRepository) ListProductsByIds(productIds []uint) (products []*model.Product, err error) {
	products = make([]*model.Product, 0, len(productIds))

	// 尝试从 Redis 缓存中获取数据
	cacheProducts, _ := p.cache.MGet(p.ctx, productIds)

	// 处理缓存结果
	missingProductIds := make([]uint, 0, len(productIds))
	for _, productId := range productIds {
		if product, exist := cacheProducts[productId]; !exist {
			missingProductIds = append(missingProductIds, productId) // 缓存未命中，添加到待查询ID列表中
		} else if exist {
			products = append(products, product) // 缓存命中，直接添加到结果中
		}
	}

	// 从数据库中获取缓存未命中的商品
	if len(missingProductIds) > 0 {
		var dbProducts []*model.Product
		// 从数据库中获取数据
		dbProducts, err = p.dao.ListByIDs(p.ctx, missingProductIds)
		if err != nil {
			return nil, err
		}

		// 将数据库查询结果批量添加到缓存中
		go p.cache.SetBatch(p.ctx, dbProducts)

		// 将数据库查询结果添加到最终结果中
		products = append(products, dbProducts...)

	}

	return products, err
}

// CreateProduct 创建商品
func (p *ProductRepository) CreateProduct(product *model.Product) (err error) {
	err = p.dao.Create(p.ctx, product)
	return
}

// DeleteProduct 删除商品
func (p *ProductRepository) DeleteProduct(productId uint) (err error) {
	if err = p.dao.Delete(p.ctx, productId); err != nil {
		return
	}

	// 删除缓存
	go p.deleteCacheWithDelay(productId)
	return

}

// UpdateProduct 更新商品
func (p *ProductRepository) UpdateProduct(product *model.Product) (err error) {
	if err = p.dao.Update(p.ctx, product); err != nil {
		return
	}

	// 删除缓存
	go p.deleteCacheWithDelay(product.ID)
	return
}

// ClearCategory 清除商品与分类的关联关系
func (p *ProductRepository) ClearCategory(product *model.Product) (err error) {
	err = p.dao.ClearCategories(p.ctx, product.ID)
	return
}

// deleteCacheWithDelay 删除缓存，延迟双删
func (p *ProductRepository) deleteCacheWithDelay(productId uint) {
	// 立即删除缓存
	_ = p.cache.Delete(context.Background(), productId)

	// 延迟双删
	time.AfterFunc(1*time.Second, func() {
		_ = p.cache.Delete(context.Background(), productId)
	})
}
