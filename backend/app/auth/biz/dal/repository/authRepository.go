package repository

import (
	"byte_go/backend/app/auth/biz/dal/redis"
	"byte_go/backend/app/auth/biz/dal/redis/dao"
	"context"
)

/**
 * @author: 锈渎
 * @date: 2025/5/8 20:47
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type AuthRepository struct {
	ctx     context.Context
	authDao *dao.AuthDao
}

func NewAuthRepository(ctx context.Context) *AuthRepository {
	return &AuthRepository{
		ctx:     ctx,
		authDao: dao.NewAuthDao(redis.RedisClient),
	}
}

func (r *AuthRepository) SetJTIBlackListed(ctx context.Context, jti string) error {
	return r.authDao.SetJTIBlackListed(ctx, jti)
}
func (r *AuthRepository) IsJTIBlackListed(ctx context.Context, jti string) (bool, error) {
	return r.authDao.IsJTIBlackListed(ctx, jti)
}
func (r *AuthRepository) GetPermissionVersion(ctx context.Context, userId uint32) (int64, error) {
	return r.authDao.GetPermissionVersion(ctx, userId)
}
func (r *AuthRepository) IncrementPermissionVersion(ctx context.Context, userId uint32) error {
	return r.authDao.IncrementPermissionVersion(ctx, userId)
}
