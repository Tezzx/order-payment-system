package repository

import (
	"order-payment-system/internal/model"

	"gorm.io/gorm"
)

type GoodsRepo struct {
	db *gorm.DB
}

func NewGoodsRepo(db *gorm.DB) *GoodsRepo {
	return &GoodsRepo{
		db: db,
	}
}

// CreateGoods 创建商品，商品名重复则覆盖旧数据
func (g *GoodsRepo) CreateGoods(goods *model.Goods) error {
	// 1. 根据商品名称查询是否已存在
	var existGoods model.Goods
	err := g.db.Where("goodsname = ?", goods.Goodsname).First(&existGoods).Error

	if err == gorm.ErrRecordNotFound {
		return g.db.Create(goods).Error
	}
	if err != nil {
		return err
	}

	return g.db.Model(&existGoods).Omit("id").Updates(goods).Error
}

// 根据商品ID获取 商品价格 和 库存数量
func (g *GoodsRepo) GetGoodsByID(goodsID uint) (price, goodsNum uint, err error) {
	var goods model.Goods
	err = g.db.Where("id = ?", goodsID).Select("price, goodsnum").First(&goods).Error
	if err != nil {
		return 0, 0, err
	}
	return goods.Price, goods.Goodsnum, nil
}
