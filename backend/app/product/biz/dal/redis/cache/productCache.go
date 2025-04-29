package cache

import (
	"byte_go/backend/app/product/biz/model"
	"byte_go/backend/utils/base-cache"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/4/27 15:00
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: product缓存

错误日志都放在base_cache中，这里只需要关注缓存的逻辑
*/

type ProductCache struct {
	base *base_cache.BaseCache[uint]
}

// NewProductCache 创建产品缓存实例
// client: Redis客户端
// prefix: 缓存键前缀
func NewProductCache(client *redis.Client) *ProductCache {
	return &ProductCache{
		base: base_cache.NewBaseCache[uint](client, base_cache.BaseCacheConfig{
			KeyPrefix:         "product:",
			DefaultExpiration: 30 * time.Minute,
			NotFoundTTL:       5 * time.Minute,
			MaxRetries:        3,
		}),
	}
}

// productUnmarshal 反序列化product数据
func productUnmarshal(data []byte) (interface{}, error) {
	var product model.Product
	if err := json.Unmarshal(data, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

// BuildProductCacheKey 构建产品缓存键
func (cache *ProductCache) BuildProductCacheKey(productId uint) string {
	return cache.base.BuildKey(productId)
}

// Get 从缓存中获取产品信息
func (cache *ProductCache) Get(ctx context.Context, productId uint) (*model.Product, error) {
	var product model.Product
	if err := cache.base.Get(ctx, productId, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

// MGet 批量获取产品信息
func (cache *ProductCache) MGet(ctx context.Context, productIds []uint) (map[uint]*model.Product, error) {
	results, err := cache.base.MGet(ctx, productIds, productUnmarshal)
	if err != nil {
		return nil, err
	}

	products := make(map[uint]*model.Product)
	for keyId, value := range results {
		if product, ok := value.(*model.Product); ok {
			products[keyId] = product
		}

	}

	return products, nil
}

// Set 缓存产品信息
func (cache *ProductCache) Set(ctx context.Context, product *model.Product) error {
	return cache.base.SafeSet(ctx, product.ID, product)
}

// SetBatch 批量缓存产品信息
func (cache *ProductCache) SetBatch(ctx context.Context, products []*model.Product) error {
	items := make(map[uint]interface{})
	for _, product := range products {
		items[product.ID] = product
	}

	return cache.base.SetBatch(ctx, items)
}

// SetNotFound 缓存空值，防止缓存穿透
func (cache *ProductCache) SetNotFound(ctx context.Context, productId uint) error {
	return cache.base.SetNotFound(ctx, productId)
}

// Delete 删除产品缓存
func (cache *ProductCache) Delete(ctx context.Context, productId uint) error {
	return cache.base.Delete(ctx, productId)
}
