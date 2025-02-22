package model

import (
	"byte_go/kitex_err"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/**
 * @author: 锈渎
 * @date: 2025/2/3 22:32
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: 用户模型
 */

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex:idx_email;type:varchar(255) not null;"`
	PasswordHash string `gorm:"type:varchar(255) not null;"`
	Username     string `gorm:"type:varchar(255) not null;"`
	Avatar       string `gorm:"type:varchar(255) not null;"`
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	// 处理唯一索引冲突
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}}, // 指定冲突列
		DoNothing: true,                             // 冲突时不操作
	}).Create(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return kitex_err.EmailExistError
	}
	return nil
}

func GetByEmail(db *gorm.DB, email string) (user User, err error) {
	if err = db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetById(db *gorm.DB, userId uint) (user User, err error) {
	if err = db.Where("id =?", userId).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DeleteById(db *gorm.DB, userId uint) error {
	return db.Where("id =?", userId).Delete(&User{}).Error
}

func UpdateById(db *gorm.DB, userId uint, updateInfo map[string]interface{}) error {

	return db.Model(&User{}).Where("id =?", userId).Updates(updateInfo).Error

}
