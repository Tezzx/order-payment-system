package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Goodsname string `gorm:"type:varchar(50);uniqueIndex"`
	Goodsnum  uint   `gorm:"type:int unsigned"`
	Price     uint   `gorm:"type:int unsigned"`
}
