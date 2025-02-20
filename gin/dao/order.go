package dao

import (
	"context"
	"gin/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) ListOrder() (Order []model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Find(&Order).Error
	return
}

func (dao *OrderDao) GetOrderById(oid uint, uid uint) (Order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", oid, uid).First(&Order).Error
	return
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (Order []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).Count(&total).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Order{}).Where(condition).Offset((page.PageNum - 1) * (page.PageNum)).Limit(page.PageSize).Find(&Order).Error
	return
}

func (dao *OrderDao) UpdateOrder(aid uint, Order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Where("id=?", aid).Updates(&Order).Error
}

func (dao *OrderDao) DeleteOrderById(uid, aid uint) error {
	return dao.DB.Model(&model.Order{}).Where("user_id=? AND id=?", uid, aid).Delete(&model.Order{}).Error
}
