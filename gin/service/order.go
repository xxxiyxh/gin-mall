package service

import (
	"context"
	"fmt"
	"gin/dao"
	"gin/model"
	"gin/pkg/e"
	"gin/serializer"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `form:"product_id" json:"product_id"`
	Num       uint    `form:"num" json:"num"`
	AddressId uint    `form:"address_id" json:"address_id"`
	Money     float64 `form:"money" json:"money"`
	BossId    uint    `form:"boss_id" json:"boss_id"`
	UserId    uint    `form:"user_id" json:"user_id"`
	OrderNum  int     `form:"order_num" json:"order_num"`
	Type      int     `form:"type" json:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uid uint) serializer.Response {
	var order *model.Order
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order = &model.Order{
		UserId:    uid,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1,
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(service.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	order.AddressId = address.ID
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.Num))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum
	err = orderDao.CreateOrder(order)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *OrderService) Show(ctx context.Context, uid uint, oid string) serializer.Response {
	orderId, _ := strconv.Atoi(oid)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderById(uint(orderId), uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(order.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uid
	orderList, total, err := orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orderList), uint(total))
}

func (service *OrderService) Delete(ctx context.Context, uid uint, aid string) serializer.Response {
	OrderId, _ := strconv.Atoi(aid)
	code := e.Success
	OrderDao := dao.NewOrderDao(ctx)
	err := OrderDao.DeleteOrderById(uid, uint(OrderId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
