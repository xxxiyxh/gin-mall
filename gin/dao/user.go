package dao

import (
	"context"
	"errors"
	"gin/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, false, nil // 如果找不到记录，返回 false
	}
	return user, err == nil, err // 如果没有错误，返回 true
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error
}

func (dao *UserDao) GetUserByID(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(user).Error
}
