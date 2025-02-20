package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	Num       uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	OrderNum  uint64
	Type      uint
	Money     float64
}
