package serializer

import (
	"context"
	"gin/dao"
	"gin/model"
)

type Cart struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	ProductId uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
	Num       int    `json:"num"`
	Name      string `json:"name"`
	MaxNum    int    `json:"max_num"`
	ImgPath   string `json:"img_path"`
	Check     bool   `json:"check"`
	Discount  string `json:"discount"`
	BossId    uint   `json:"boss_id"`
	BossName  string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		Id:        cart.ID,
		UserId:    cart.UserId,
		ProductId: product.ID,
		CreateAt:  cart.CreatedAt.Unix(),
		Num:       int(cart.Num),
		MaxNum:    int(cart.MaxNum),
		Check:     cart.Check,
		Name:      product.Name,
		ImgPath:   product.ImgPath,
		Discount:  product.DiscountPrice,
		BossId:    boss.ID,
		BossName:  boss.UserName,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []Cart) {
	cartDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := cartDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserByID(item.BossId)
		if err != nil {
			continue
		}
		cart := BuildCart(item, product, boss)
		carts = append(carts, cart)
	}
	return carts
}
