package dao

import (
	"byte_go/backend/app/user/biz/model"
	"byte_go/kitex_err"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/**
 * @author: 锈渎
 * @date: 2025/4/28 15:56
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Create(ctx context.Context, user *model.User) error {
	// 处理唯一索引冲突
	result := dao.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}}, // 指定冲突列
		DoNothing: true,                             // 冲突时不操作
	}).Create(user)

	if result.Error != nil {
		klog.Errorf("mysql: create user failed: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		klog.Errorf("mysql: create user failed: %v", kitex_err.EmailExistError)
		return kitex_err.EmailExistError
	}
	return nil
}

func (dao *UserDAO) GetByEmail(ctx context.Context, email string) (user *model.User, err error) {
	if err = dao.db.WithContext(ctx).Where("email =?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			klog.Infof("mysql: get user by email [%s] not found", email)
			return nil, kitex_err.UserNotExist
		}
		klog.Errorf("mysql: get user by email failed: %v", err.Error())
		return nil, err
	}
	return user, nil
}

func (dao *UserDAO) GetById(ctx context.Context, userId uint) (user *model.User, err error) {
	if err = dao.db.WithContext(ctx).Where("id =?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			klog.Infof("mysql: get user by id [%d] not found", userId)
			return nil, kitex_err.UserNotExist
		}
		klog.Errorf("mysql: get user by id failed: %v", err.Error())
		return nil, err
	}
	return user, nil
}

func (dao *UserDAO) DeleteById(ctx context.Context, userId uint) error {
	if err := dao.db.WithContext(ctx).Where("id =?", userId).Delete(&model.User{}).Error; err != nil {
		klog.Errorf("mysql: delete user by id failed: %v", err.Error())
		return err
	}
	return nil
}

func (dao *UserDAO) Update(ctx context.Context, userId uint, updateInfo map[string]interface{}) error {
	result := dao.db.WithContext(ctx).Model(&model.User{}).Where("id =?", userId).Updates(updateInfo)
	if result.Error != nil {
		klog.Errorf("mysql: update user failed: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		klog.Errorf("mysql: update user failed: %v", kitex_err.UserNotExist)
		return kitex_err.UserNotExist
	}
	return nil
}
