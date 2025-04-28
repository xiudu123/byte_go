package cache

import (
	"byte_go/backend/app/product/biz/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/4/27 15:00
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type ProductCache struct {
	client *redis.Client
	prefix string
}

// NewProductCache 创建产品缓存实例
// client: Redis客户端
// prefix: 缓存键前缀
func NewProductCache(client *redis.Client) *ProductCache {
	return &ProductCache{
		client: client,
		prefix: ProductCacheKeyPrefix,
	}
}

// BuildProductCacheKey 构建产品缓存键
func (cache *ProductCache) BuildProductCacheKey(productId uint) string {
	return fmt.Sprintf("%s%d", cache.prefix, productId)
}

// safeSetCache 带重试机制的缓存写入
func (cache *ProductCache) safeSetCache(ctx context.Context, key string, value []byte) error {
	const maxRetries = 3
	var err error
	for i := 0; i < maxRetries; i++ {
		if err = cache.client.Set(ctx, key, value, randomExpiration(ProductCacheExpire)).Err(); err == nil {
			return nil
		}
		time.Sleep(time.Duration(i*100) * time.Millisecond)
	}
	return err
}

// Get 从缓存中获取产品信息
func (cache *ProductCache) Get(ctx context.Context, productId uint) (*model.Product, error) {
	cacheKey := cache.BuildProductCacheKey(productId)

	data, err := cache.client.Get(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrCacheMiss
		}
		klog.Errorf("redis: get product cache failed: %v", err.Error())
		return nil, err
	}
	if data == NotFoundMarker {
		klog.Warnf("redis: product by id [%d] is null value", productId)
		return &model.Product{}, nil
	}

	var product model.Product
	if err := json.Unmarshal([]byte(data), &product); err != nil {
		klog.Errorf("redis: unmarshal product cache failed: %v", err.Error())
		return nil, ErrInvalidCache
	}

	return &product, nil
}

// MGet 批量获取产品信息
func (cache *ProductCache) MGet(ctx context.Context, productIds []uint) (map[uint]*model.Product, error) {
	cacheKeys := make([]string, len(productIds))

	for i, productId := range productIds {
		cacheKeys[i] = cache.BuildProductCacheKey(productId)
	}

	results, err := cache.client.MGet(ctx, cacheKeys...).Result()
	if err != nil {
		klog.Errorf("redis: mGet product cache failed: %v", err.Error())
		return nil, err
	}

	products := make(map[uint]*model.Product)
	for idx, result := range results {
		if result == nil {
			continue
		}
		if result == NotFoundMarker {
			products[productIds[idx]] = nil
			continue
		}
		data, ok := result.(string)
		if !ok {
			continue
		}
		var product model.Product
		if err := json.Unmarshal([]byte(data), &product); err == nil {
			products[productIds[idx]] = &product
		}
	}

	return products, nil
}

// Set 缓存产品信息
func (cache *ProductCache) Set(ctx context.Context, product *model.Product) error {
	cacheKey := cache.BuildProductCacheKey(product.ID)

	data, err := json.Marshal(product)

	if err != nil {
		klog.Errorf("redis: marshal product cache failed: %v", err.Error())
		return err
	}

	if err = cache.safeSetCache(ctx, cacheKey, data); err != nil {
		klog.Errorf("redis: set product cache failed: %v", err.Error())
		return err
	}

	return nil
}

// SetBatch 批量缓存产品信息
func (cache *ProductCache) SetBatch(ctx context.Context, products []*model.Product) {
	// 使用管道批量写入缓存
	pipe := cache.client.Pipeline()
	for _, product := range products {
		cacheKey := cache.BuildProductCacheKey(product.ID)
		if data, err := json.Marshal(product); err == nil {
			pipe.Set(ctx, cacheKey, data, randomExpiration(ProductCacheExpire))
		}
	}

	if _, err := pipe.Exec(ctx); err != nil {
		klog.Errorf("redis pipe set product cache failed: %v", err.Error())
	}
}

// SetNotFound 缓存空值，防止缓存穿透
func (cache *ProductCache) SetNotFound(ctx context.Context, productId uint) error {
	cacheKey := cache.BuildProductCacheKey(productId)
	err := cache.client.Set(ctx, cacheKey, NotFoundMarker, NotFoundTTL).Err()

	if err != nil {
		klog.Errorf("redis: set product not found cache failed: %v", err.Error())
	}

	return err
}

// Delete 删除产品缓存
func (cache *ProductCache) Delete(ctx context.Context, productId uint) error {
	cacheKey := cache.BuildProductCacheKey(productId)
	err := cache.client.Del(ctx, cacheKey).Err()

	if err != nil {
		klog.Errorf("redis: delete product cache failed: %v", err.Error())
	}

	return err
}
