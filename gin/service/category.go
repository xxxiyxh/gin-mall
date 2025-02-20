package service

import (
	"context"
	"gin/dao"
	"gin/pkg/e"
	"gin/pkg/util"
	"gin/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	CategoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	Categorys, err := CategoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategorys(Categorys), uint(len(Categorys)))
}
