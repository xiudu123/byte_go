package cache

import (
	"byte_go/backend/app/user/biz/model"
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
 * @date: 2025/4/28 16:09
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

const (
	UserByIdCacheKey    = "user:id:"
	UserByEmailCacheKey = "user:email:"
	DefaultExpiration   = 24 * time.Hour
	NotFoundTTL         = 5 * time.Minute
	NotFoundMarker      = "NOT_FOUND"
)

var (
	ErrCacheMiss    = errors.New("cache miss")
	ErrInvalidCache = errors.New("invalid cache")
)

type UserCache struct {
	client      *redis.Client
	prefixId    string
	prefixEmail string
}

func randomExpiration(cacheExpire time.Duration) time.Duration {
	return cacheExpire + time.Duration(rand.Intn(300))*time.Second
}

func NewUserCache(client *redis.Client) *UserCache {
	return &UserCache{
		client:      client,
		prefixId:    UserByIdCacheKey,
		prefixEmail: UserByEmailCacheKey,
	}
}
func (cache *UserCache) safeSetCache(ctx context.Context, key string, value []byte) error {
	const maxRetries = 3
	var err error
	for i := 0; i < maxRetries; i++ {
		if err = cache.client.Set(ctx, key, value, randomExpiration(DefaultExpiration)).Err(); err == nil {
			return nil
		}
		time.Sleep(time.Duration(i*100) * time.Millisecond)
	}
	return err
}

func (cache *UserCache) BuildUserIdKey(userId uint) string {
	return fmt.Sprintf("%s%d", cache.prefixId, userId)
}

func (cache *UserCache) BuildUserEmailKey(email string) string {
	return fmt.Sprintf("%s%s", cache.prefixEmail, email)
}

func (cache *UserCache) GetById(ctx context.Context, userId uint) (*model.User, error) {
	cacheKey := cache.BuildUserIdKey(userId)
	data, err := cache.client.Get(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrCacheMiss
		}
		klog.Errorf("redis: get user by id failed: %v", err.Error())
		return nil, err
	}
	if data == NotFoundMarker {
		klog.Warnf("redis: user by id [%d] is not value", userId)
		return &model.User{}, nil
	}
	var user model.User
	if err = json.Unmarshal([]byte(data), &user); err != nil {
		klog.Errorf("redis: unmarshal user by id failed: %v", err.Error())
		return nil, ErrInvalidCache
	}
	return &user, nil
}

func (cache *UserCache) GetByEmail(ctx context.Context, email string) (userId uint, err error) {
	cacheKey := cache.BuildUserEmailKey(email)
	data, err := cache.client.Get(ctx, cacheKey).Result()

	if err != nil {
		if err == redis.Nil {
			return 0, ErrCacheMiss
		}
		klog.Errorf("redis: get user by email failed: %v", err.Error())
		return 0, err
	}
	if data == NotFoundMarker {
		klog.Warnf("redis: user by email [%s] is not value", email)
		return 0, nil
	}

	if err = json.Unmarshal([]byte(data), &userId); err != nil {
		klog.Errorf("redis: unmarshal user by email failed: %v", err.Error())
		return 0, ErrInvalidCache
	}
	return userId, nil
}

func (cache *UserCache) Set(ctx context.Context, user *model.User) error {
	cacheKey := cache.BuildUserIdKey(user.ID)
	userData, err := json.Marshal(user)
	if err != nil {
		klog.Errorf("redis: marshal user failed: %v", err.Error())
		return err
	}

	// 缓存 邮箱 -> ID 映射
	emailKey := cache.BuildUserEmailKey(user.Email)
	pipe := cache.client.Pipeline()
	pipe.Set(ctx, cacheKey, userData, randomExpiration(DefaultExpiration))
	pipe.Set(ctx, emailKey, user.ID, randomExpiration(DefaultExpiration))
	if _, err := pipe.Exec(ctx); err != nil {
		klog.Errorf("redis: set user failed: %v", err.Error())
		return err
	}
	return nil
}
func (cache *UserCache) Delete(ctx context.Context, user *model.User) error {
	cacheKeys := []string{
		cache.BuildUserIdKey(user.ID),
		cache.BuildUserEmailKey(user.Email),
	}
	if _, err := cache.client.Del(ctx, cacheKeys...).Result(); err != nil {
		klog.Errorf("redis: delete user failed: %v", err.Error())
		return err
	}
	return nil
}

func (cache *UserCache) SetNotFound(ctx context.Context, keyType string, value interface{}) error {
	var cacheKey string
	switch keyType {
	case UserByIdCacheKey:
		cacheKey = cache.BuildUserIdKey(value.(uint))
	case UserByEmailCacheKey:
		cacheKey = cache.BuildUserEmailKey(value.(string))
	default:
		return fmt.Errorf("invalid key type: %s", keyType)
	}
	return cache.client.Set(ctx, cacheKey, NotFoundMarker, NotFoundTTL).Err()
}
