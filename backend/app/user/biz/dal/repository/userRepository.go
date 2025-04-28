package repository

import (
	"byte_go/backend/app/user/biz/dal/mysql"
	"byte_go/backend/app/user/biz/dal/mysql/dao"
	"byte_go/backend/app/user/biz/dal/redis"
	"byte_go/backend/app/user/biz/dal/redis/cache"
	"byte_go/backend/app/user/biz/model"
	"byte_go/kitex_err"
	"context"
	"errors"
	"golang.org/x/sync/singleflight"
)

/**
 * @author: 锈渎
 * @date: 2025/4/28 16:45
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type UserRepository struct {
	ctx   context.Context
	dao   *dao.UserDAO
	cache *cache.UserCache
	sg    *singleflight.Group
}

func NewUserRepository(ctx context.Context) *UserRepository {
	return &UserRepository{
		ctx:   ctx,
		dao:   dao.NewUserDAO(mysql.DB),
		cache: cache.NewUserCache(redis.RedisClient),
		sg:    &singleflight.Group{},
	}
}

func (r *UserRepository) GetById(userId uint) (*model.User, error) {
	// 从缓存中获取数据
	if cacheUser, err := r.cache.GetById(r.ctx, userId); err == nil {
		return cacheUser, nil
	}

	// 使用SingleFlight模式，确保同一时间只有一个请求查询数据库, 防止缓存击穿
	result, err, _ := r.sg.Do("get user from mysql "+r.cache.BuildUserIdKey(userId), func() (interface{}, error) {
		// 从数据库查询数据
		dbUser, err := r.dao.GetById(r.ctx, userId)
		if err != nil {
			// 数据库不存在数据
			if errors.Is(err, kitex_err.UserNotExist) {
				go func() {
					_ = r.cache.SetNotFound(context.Background(), cache.UserByIdCacheKey, userId) // 缓存空值，防止缓存穿透
				}()
				return &model.User{}, nil
			}
			return nil, err
		}

		go func() {
			_ = r.cache.Set(context.Background(), dbUser)
		}()

		return dbUser, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*model.User), nil
}
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	userId, err := r.cache.GetByEmail(r.ctx, email)
	if err == nil {
		if userId == 0 {
			return &model.User{}, nil
		}
		return r.GetById(userId)
	}

	user, err := r.dao.GetByEmail(r.ctx, email)
	if err != nil {
		if errors.Is(err, kitex_err.UserNotExist) {
			go func() {
				_ = r.cache.SetNotFound(context.Background(), cache.UserByEmailCacheKey, email) // 缓存空值，防止缓存穿透
			}()
			return &model.User{}, nil
		}
		return nil, err
	}

	go func() {
		_ = r.cache.Set(context.Background(), user)
	}()

	return user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	return r.dao.Create(r.ctx, user)
}

func (r *UserRepository) UpdateById(userId uint, updateInfo map[string]interface{}) error {
	oldUser, err := r.GetById(userId)
	if err != nil {
		return err
	}

	if err := r.dao.Update(r.ctx, userId, updateInfo); err != nil {
		return err
	}
	go func() {
		_ = r.cache.Delete(context.Background(), oldUser)
	}()
	return nil
}

func (r *UserRepository) DeleteById(userId uint) error {
	user, err := r.GetById(userId)
	if err != nil {
		return err
	}
	if err := r.dao.DeleteById(r.ctx, userId); err != nil {
		return err
	}
	go func() {
		_ = r.cache.Delete(context.Background(), user)
	}()
	return nil
}
