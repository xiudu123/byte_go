package dao

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/5/8 20:31
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

const (
	JTIBlackListedKey    = "jti:blacklisted:"
	PermissionVersionKey = "permission:version:"
	DefaultExpiration    = 24 * time.Hour
)

type AuthDao struct {
	client *redis.Client
}

func NewAuthDao(client *redis.Client) *AuthDao {
	return &AuthDao{
		client: client,
	}
}

func randomExpiration() time.Duration {
	return DefaultExpiration + time.Duration(rand.Intn(300))*time.Second
}

func (dao *AuthDao) SetJTIBlackListed(ctx context.Context, jti string) error {
	err := dao.client.Set(ctx, JTIBlackListedKey+jti, true, randomExpiration()).Err()
	if err != nil {
		klog.Errorf("redis auth dao: set jti blacklisted failed: %v", err.Error())
		return err
	}
	return nil
}

func (dao *AuthDao) IsJTIBlackListed(ctx context.Context, jti string) (bool, error) {
	exists, err := dao.client.Exists(ctx, JTIBlackListedKey+jti).Result()
	if err != nil {
		klog.Errorf("redis auth dao: check jti blacklisted failed: %v", err.Error())
		return false, err
	}
	return exists == 1, nil
}

func (dao *AuthDao) GetPermissionVersion(ctx context.Context, userId uint32) (int64, error) {
	key := fmt.Sprintf("%s%d", PermissionVersionKey, userId)
	version, err := dao.client.Get(ctx, key).Int64()
	if err != nil {
		klog.Errorf("redis auth dao: get permission version failed: %v", err.Error())
		return 0, err
	}
	return version, nil
}

func (dao *AuthDao) IncrementPermissionVersion(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", PermissionVersionKey, userId)

	if err := dao.client.Incr(ctx, key).Err(); err != nil {
		klog.Errorf("redis auth dao: increment permission version failed: %v", err.Error())
		return err
	}

	if err := dao.client.Expire(ctx, key, randomExpiration()).Err(); err != nil {
		klog.Errorf("redis auth dao: set permission version expiration failed: %v", err.Error())
		return err
	}

	return nil
}
