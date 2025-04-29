package cache

import (
	"byte_go/backend/app/user/biz/model"
	base_cache "byte_go/backend/utils/base-cache"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/4/28 16:09
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: user缓存

错误日志都放在base_cache中，这里只需要关注缓存的逻辑
*/

const (
	UserByIdCacheKey    = "user:id"
	UserByEmailCacheKey = "user:email"
	DefaultExpiration   = 24 * time.Hour
	NotFoundTTL         = 5 * time.Minute
)

type UserCache struct {
	base      *base_cache.BaseCache[uint]
	emailPool *base_cache.BaseCache[string]
}

func NewUserCache(client *redis.Client) *UserCache {
	return &UserCache{
		base: base_cache.NewBaseCache[uint](client, base_cache.BaseCacheConfig{
			KeyPrefix:         UserByIdCacheKey,
			DefaultExpiration: DefaultExpiration,
			NotFoundTTL:       NotFoundTTL,
			MaxRetries:        3,
		}),
		emailPool: base_cache.NewBaseCache[string](client, base_cache.BaseCacheConfig{
			KeyPrefix:         UserByEmailCacheKey,
			DefaultExpiration: DefaultExpiration,
			NotFoundTTL:       NotFoundTTL,
			MaxRetries:        3,
		}),
	}
}

// userUnmarshal 反序列化user数据
func userUnmarshal(data []byte) (interface{}, error) {
	var u model.User
	if err := json.Unmarshal(data, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

// idUnmarshal 反序列化id数据(用于邮箱映射)
func idUnmarshal(data []byte) (interface{}, error) {
	var id uint
	if err := json.Unmarshal(data, &id); err != nil {
		return nil, err
	}
	return id, nil
}

// BuildUserIdKey 构建用户ID缓存键
func (cache *UserCache) BuildUserIdKey(userId uint) string {
	return cache.base.BuildKey(userId)
}

// BuildUserEmailKey 构建用户邮箱缓存键
func (cache *UserCache) BuildUserEmailKey(email string) string {
	return cache.emailPool.BuildKey(email)
}

// GetById 从缓存中获取用户信息
func (cache *UserCache) GetById(ctx context.Context, userId uint) (*model.User, error) {
	var user model.User
	if err := cache.base.Get(ctx, userId, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail 从缓存中获取用户ID
func (cache *UserCache) GetByEmail(ctx context.Context, email string) (userId uint, err error) {
	if err = cache.emailPool.Get(ctx, email, &userId); err != nil {
		return 0, err
	}
	return userId, nil
}

// Set 缓存用户信息
func (cache *UserCache) Set(ctx context.Context, user *model.User) error {

	// 创建带取消功能的错误组
	g, ctx := errgroup.WithContext(ctx)

	// 并发设置缓存
	g.Go(func() error {
		return cache.base.SafeSet(ctx, user.ID, user)
	})
	g.Go(func() error {
		return cache.emailPool.SafeSet(ctx, user.Email, user.ID)
	})

	// 等待所有操作完成
	return g.Wait()
}

// Delete 删除用户缓存
func (cache *UserCache) Delete(ctx context.Context, user *model.User) error {

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return cache.base.Delete(ctx, user.ID)
	})

	g.Go(func() error {
		return cache.emailPool.Delete(ctx, user.Email)
	})

	return g.Wait()
}

// SetNotFound 设置空值缓存
func (cache *UserCache) SetNotFound(ctx context.Context, keyType string, value interface{}) error {
	switch keyType {
	case UserByIdCacheKey:
		return cache.base.SetNotFound(ctx, value.(uint))
	case UserByEmailCacheKey:
		return cache.emailPool.SetNotFound(ctx, value.(string))
	default:
		return fmt.Errorf("invalid key type: %s", keyType)
	}
}
