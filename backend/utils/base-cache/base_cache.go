package base_cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/4/28 21:04
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: 缓存基础操作
 */

var (
	ErrCacheMiss    = errors.New("base-cache miss")
	ErrInvalidCache = errors.New("invalid base-cache")
	ErrNullValue    = errors.New("null value")
	ErrCacheSystem  = errors.New("base-cache system error")
	ErrCanceled     = errors.New("base-cache canceled")
	NoValueMarker   = ""
)

type keyType interface {
	~uint | ~uint32 | ~uint64 | ~int | ~int32 | ~int64 | ~string
}

type BaseCacheConfig struct {
	KeyPrefix         string        // 缓存建前缀
	DefaultExpiration time.Duration // 默认缓存时间
	NotFoundTTL       time.Duration // 空值缓存时间
	MaxRetries        int           // 最大重试次数
}

type BaseCache[T keyType] struct {
	client *redis.Client
	config BaseCacheConfig
}

func NewBaseCache[T keyType](client *redis.Client, config BaseCacheConfig) *BaseCache[T] {
	if config.NotFoundTTL == 0 {
		config.NotFoundTTL = 5 * time.Minute
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}
	if config.DefaultExpiration == 0 {
		config.DefaultExpiration = 24 * time.Hour
	}
	return &BaseCache[T]{
		client: client,
		config: config,
	}
}

// randomExpiration 生成随机过期时间
func (cache *BaseCache[T]) randomExpiration() time.Duration {
	return cache.config.DefaultExpiration + time.Duration(rand.Intn(300))*time.Second
}

// BuildKey 构建缓存键
func (cache *BaseCache[T]) BuildKey(keyId T) string {
	return fmt.Sprintf("%s%v", cache.config.KeyPrefix, keyId)
}

// BuildKeys 批量构建缓存键
func (cache *BaseCache[T]) BuildKeys(keyIds []T) []string {
	keys := make([]string, len(keyIds))
	for i, keyId := range keyIds {
		keys[i] = cache.BuildKey(keyId)
	}
	return keys
}

// SafeSet 带重试机制的缓存写入
func (cache *BaseCache[T]) SafeSet(ctx context.Context, keyId T, value interface{}) error {
	key := cache.BuildKey(keyId)
	// 序列化数据
	data, err := json.Marshal(value)
	if err != nil {
		klog.Errorf("redis base base-cache: marshal failed: %v", err.Error())
		return ErrCacheSystem
	}

	// 重试机制
	for i := 0; i < cache.config.MaxRetries; i++ {
		if err = cache.client.Set(ctx, key, data, cache.randomExpiration()).Err(); err == nil {
			return nil
		}
		time.Sleep(time.Duration(i*100) * time.Millisecond)
	}

	select {
	case <-ctx.Done():
		klog.Errorf("redis base base-cache: set base-cache canceled: %v", ctx.Err().Error())
		return ErrCanceled
	default:
		klog.Errorf("redis base base-cache: set base-cache failed after %d retries: %v", cache.config.MaxRetries, err.Error())
		return ErrCacheSystem
	}
}

// Get 从缓存中获取数据
func (cache *BaseCache[T]) Get(ctx context.Context, keyId T, result interface{}) error {
	key := cache.BuildKey(keyId)
	// 从缓存中获取数据
	data, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return ErrCacheMiss
		}
		klog.Errorf("redis base base-cache: get base-cache failed: %v", err.Error())
		return ErrCacheSystem
	}

	// 处理空值情况
	if data == NoValueMarker {
		klog.Warnf("redis base base-cache: base-cache by key [%s] is null value", key)
		return ErrNullValue
	}

	// 解析缓存数据
	if err = json.Unmarshal([]byte(data), result); err != nil {
		klog.Errorf("redis base base-cache: unmarshal base-cache failed: %v", err.Error())
		return ErrInvalidCache
	}
	return nil
}

// MGet 批量获取缓存数据
func (cache *BaseCache[T]) MGet(ctx context.Context, keyIds []T,
	unmarshalFuc func([]byte) (interface{}, error)) (map[T]interface{}, error) {
	data := make(map[T]interface{})

	if len(keyIds) == 0 {
		return data, nil
	}

	// 构建缓存键列表
	keys := cache.BuildKeys(keyIds)

	// 从缓存中批量获取数据
	results, err := cache.client.MGet(ctx, keys...).Result()
	if err != nil {
		klog.Errorf("redis base base-cache: mget base-cache failed: %v", err.Error())
		return nil, ErrCacheSystem
	}

	// 解析缓存数据
	for i, key := range keys {
		if results[i] == nil {
			continue
		}
		strVal, ok := results[i].(string)
		if !ok {
			klog.Errorf("redis base base-cache: invalid base-cache type for key [%s]", key)
			continue
		}
		if strVal == NoValueMarker {
			klog.Warnf("redis base base-cache: base-cache by key [%s] is null value", key)
			continue
		}
		value, err := unmarshalFuc([]byte(strVal))
		if err != nil {
			klog.Errorf("redis base base-cache: unmarshal base-cache failed for key [%s]: %v", key, err.Error())
			continue
		}
		data[keyIds[i]] = value
	}

	return data, nil
}

// SetBatch 批量设置缓存
func (cache *BaseCache[T]) SetBatch(ctx context.Context, items map[T]interface{}) error {

	pipe := cache.client.Pipeline()

	for keyId, value := range items {
		key := cache.BuildKey(keyId)
		data, err := json.Marshal(value)
		if err != nil {
			klog.Errorf("redis base base-cache: marshal failed for key [%s]: %v", key, err.Error())
			continue
		}
		pipe.Set(ctx, key, data, cache.randomExpiration())
	}

	// 执行管道
	if _, err := pipe.Exec(ctx); err != nil {
		klog.Errorf("redis base base-cache: set batch base-cache failed: %v", err.Error())
		return ErrCacheSystem
	}
	return nil
}

// SetNotFound 设置空值缓存
func (cache *BaseCache[T]) SetNotFound(ctx context.Context, keyId T) error {
	key := cache.BuildKey(keyId)

	// 设置空值缓存
	if err := cache.client.Set(ctx, key, NoValueMarker, cache.config.NotFoundTTL).Err(); err != nil {
		klog.Errorf("redis base base-cache: set not found base-cache [%s] failed: %v", key, err.Error())
		return ErrCacheSystem
	}
	return nil
}

// Delete 删除缓存
func (cache *BaseCache[T]) Delete(ctx context.Context, keyIds ...T) error {
	if len(keyIds) == 0 {
		return nil
	}

	// 构建缓存键列表
	keys := cache.BuildKeys(keyIds)

	// 执行删除操作
	if _, err := cache.client.Del(ctx, keys...).Result(); err != nil {
		klog.Errorf("redis base base-cache: delete base-cache [%v] failed: %v", keys, err.Error())
		return ErrCacheSystem
	}
	return nil
}
