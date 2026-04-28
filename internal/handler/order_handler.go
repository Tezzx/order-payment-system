package handler

import (
	"order-payment-system/internal/service"
	"order-payment-system/pkg/response"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *service.OrderService
}

type OrderRequest struct {
	GoodsID int `json:"goodsId"`
	BuyNum  int `json:"buyNum"`
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// 创建订单
func (o *OrderHandler) CreateOrder(c *gin.Context) {

	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}

	userIDany, bol := c.Get("userID")

	if !bol {
		c.JSON(200, gin.H{"code": 401, "msg": "请先登录"})
		return
	}
	userID, ok := userIDany.(uint)
	if !ok {
		c.JSON(200, gin.H{"code": 401, "msg": "登录信息无效"})
		return
	}
	//创建订单
	order, err := o.orderService.CreateOrder(userID, uint(req.GoodsID), uint(req.BuyNum))
	if err != nil {
		response.Error(c, 500, "订单创建失败")
		return
	}

	// 4. 返回成功响应
	response.Success(c, order.OrderNo)
}

func (o *OrderHandler) ToPay(c *gin.Context) {
	c.HTML(200, "pay.html", nil)

}
