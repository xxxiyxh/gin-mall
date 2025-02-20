package serializer

import (
	"gin/conf"
	"gin/model"
)

type ProductImg struct {
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductID: item.ProductId,
		ImgPath:   conf.Host + conf.HttpPort + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImg []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImg = append(productImg, product)
	}
	return
}
