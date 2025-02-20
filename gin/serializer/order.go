package serializer

import (
	"context"
	"gin/dao"
	"gin/model"
)

type Order struct {
	Id           uint    `json:"id"`
	OrderNum     uint64  `json:"order_num"`
	CreateAt     int64   `json:"create_at"`
	UpdateAt     int64   `json:"update_at"`
	UserId       uint    `json:"user_id"`
	ProductId    uint    `json:"product_id"`
	BossId       uint    `json:"boss_id"`
	Num          uint    `json:"num"`
	AddressName  string  `json:"address_name"`
	AddressPhone string  `json:"address_phone"`
	Address      string  `json:"address"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"product_name"`
	ImgPath      string  `json:"img_path"`
	Money        float64 `json:"money"`
}

func BuildOrder(item *model.Order, product *model.Product, address *model.Address) Order {
	return Order{
		Id:           item.ID,
		OrderNum:     item.OrderNum,
		CreateAt:     item.CreatedAt.Unix(),
		UpdateAt:     item.UpdatedAt.Unix(),
		UserId:       item.ID,
		ProductId:    product.ID,
		BossId:       item.ID,
		Num:          item.Num,
		AddressName:  address.Name,
		AddressPhone: address.Phone,
		Address:      address.Phone,
		Type:         item.Type,
		ProductName:  product.Name,
		ImgPath:      product.ImgPath,
		Money:        item.Money,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (orders []Order) {
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		address, err := addressDao.GetAddressById(item.BossId)
		if err != nil {
			continue
		}
		order := BuildOrder(item, product, address)
		orders = append(orders, order)
	}
	return orders
}
