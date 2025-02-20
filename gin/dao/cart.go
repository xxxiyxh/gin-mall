package dao

import (
	"context"
	"gin/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) ListCart() (Cart []model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Find(&Cart).Error
	return
}

func (dao *CartDao) GetCartById(aid uint) (Cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=?", aid).First(&Cart).Error
	return
}

func (dao *CartDao) ListCartByUid(uid uint) (Cart []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", uid).First(&Cart).Error
	return
}

func (dao *CartDao) UpdateCartNum(aid uint, num int) error {
	return dao.DB.Model(&model.Cart{}).Where("id=?", aid).Update("num", num).Error
}

func (dao *CartDao) DeleteCartById(uid, aid uint) error {
	return dao.DB.Model(&model.Cart{}).Where("user_id=? AND id=?", uid, aid).Delete(&model.Cart{}).Error
}
