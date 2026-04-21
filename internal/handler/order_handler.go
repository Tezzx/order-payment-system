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

//获取我的订单
/*
func (o *OrderHandler) GetMyOrderList(c *gin.Context) {
	userID := uint(1) // 正式环境解析Token获取
	orders, err := o.orderService.GetUserOrderList(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "获取订单失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": orders})
}
*/

//支付订单接口
/*
func (o *OrderHandler) PayOrder(c *gin.Context) {
	orderIDStr := c.PostForm("orderId")
	orderID, _ := strconv.Atoi(orderIDStr)

	err := o.orderService.PayOrder(uint(orderID))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "支付失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "支付成功"})
}
*/
