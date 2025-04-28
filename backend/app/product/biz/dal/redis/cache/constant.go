package cache

import (
	"errors"
	"math/rand"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/4/27 21:20
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */
const (
	ProductCacheKeyPrefix = "product:"
	ProductCacheExpire    = 30 * time.Minute
	NotFoundMarker        = "NOT_FOUND"
	NotFoundTTL           = 5 * time.Minute
)

var (
	ErrCacheMiss    = errors.New("cache miss")    // 缓存未命中
	ErrInvalidCache = errors.New("invalid cache") // 数据解析错误
)

// randomExpiration 生成随机过期时间
func randomExpiration(cacheExpire time.Duration) time.Duration {
	return time.Duration(rand.Intn(300))*time.Second + cacheExpire
}
