package service

import (
	"context"
	"gin/dao"
	"gin/model"
	"gin/pkg/e"
	"gin/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uid uint) serializer.Response {
	var Cart model.Cart
	code := e.Success
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	CartDao := dao.NewCartDao(ctx)
	cart := &model.Cart{
		UserId:    uid,
		ProductId: service.ProductId,
		BossId:    service.BossId,
	}
	err = CartDao.CreateCart(&Cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserByID(service.BossId)
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
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	CartDao := dao.NewCartDao(ctx)
	Cart, err := CartDao.ListCartByUid(uid)
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
		Data:   serializer.BuildCarts(ctx, Cart),
	}
}

func (service *CartService) Update(ctx context.Context, uid uint, cid string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(cid)
	err := cartDao.UpdateCartNum(uint(cartId), service.Num)
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

func (service *CartService) Delete(ctx context.Context, uid uint, aid string) serializer.Response {
	CartId, _ := strconv.Atoi(aid)
	code := e.Success
	CartDao := dao.NewCartDao(ctx)
	err := CartDao.DeleteCartById(uid, uint(CartId))
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
