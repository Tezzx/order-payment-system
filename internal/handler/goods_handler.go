package handler

import (
	"log"
	"order-payment-system/internal/model"
	"order-payment-system/internal/service"
)

type GoodsHandler struct {
	goodsService *service.GoodsService
}

func NewGoodsHandler(goodsService *service.GoodsService) *GoodsHandler {
	return &GoodsHandler{
		goodsService: goodsService,
	}
}

// 目前没有后台管理部分
// 初始化商品清单
func (g *GoodsHandler) GoodsInitial() {
	goods_01 := model.Goods{
		Goodsname: "雪影娃娃",
		Goodsnum:  10,
		Price:     1200,
	}
	goods_02 := model.Goods{
		Goodsname: "恶魔狼",
		Goodsnum:  15,
		Price:     600,
	}
	goods_03 := model.Goods{
		Goodsname: "治愈兔",
		Goodsnum:  10,
		Price:     1800,
	}
	goods_04 := model.Goods{
		Goodsname: "月牙雪熊",
		Goodsnum:  5,
		Price:     1800,
	}
	// 核心：接收错误并打印！！！
	log.Println("开始初始化商品...")

	// 批量插入并打印错误
	if err := g.goodsService.CreateGoods(&goods_01); err != nil {
		log.Fatalf("商品1插入失败：%v", err) // Fatal会打印错误并停止程序
	}
	if err := g.goodsService.CreateGoods(&goods_02); err != nil {
		log.Fatalf("商品2插入失败：%v", err)
	}
	if err := g.goodsService.CreateGoods(&goods_03); err != nil {
		log.Fatalf("商品3插入失败：%v", err)
	}
	if err := g.goodsService.CreateGoods(&goods_04); err != nil {
		log.Fatalf("商品4插入失败：%v", err)
	}

	log.Println("✅ 所有商品初始化成功！")

}
