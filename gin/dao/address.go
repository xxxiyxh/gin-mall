package dao

import (
	"context"
	"gin/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) CreateAddress(in *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&in).Error
}

func (dao *AddressDao) ListAddress() (Address []model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Find(&Address).Error
	return
}

func (dao *AddressDao) GetAddressById(aid uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aid).First(&address).Error
	return
}

func (dao *AddressDao) ListAddressByUid(uid uint) (address []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id=?", uid).First(&address).Error
	return
}

func (dao *AddressDao) UpdateAddress(aid uint, address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where("id=?", aid).Updates(&address).Error
}

func (dao *AddressDao) DeleteAddressById(uid, aid uint) error {
	return dao.DB.Model(&model.Address{}).Where("user_id=? AND id=?", uid, aid).Delete(&model.Address{}).Error
}
