package model

import (
	"gorm.io/gorm"
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
	return db.Create(user).Error
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