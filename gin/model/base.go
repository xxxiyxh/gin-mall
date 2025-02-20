package model

type BasePage struct {
	PageNum  int `form:"page_Num"`
	PageSize int `form:"page_Size"`
}
