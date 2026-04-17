package handler

import (
	"order-payment-system/internal/service"
	"order-payment-system/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

type LoginRequest struct {
	Username string
	Password string
}

type RegisterRequest struct {
	Username string
	Password string
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// 注册
func (u *UserHandler) RegisterUser(c *gin.Context) {
	var req RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Error(c, "参数错误")
		return
	}
	err = u.userService.RegisterUser(req.Username, req.Password)
	if err != nil {
		response.Error(c, "新用户创建失败")
		return
	}
	response.Success(c, "注册成功")
}

// 登录
func (u *UserHandler) LoginUser(c *gin.Context) {
	var req LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Error(c, "参数错误")
		return
	}
	s := u.userService.LoginUser(req.Username, req.Password)
	if s != "" {
		response.Error(c, s)
		return
	}
	response.Success(c, "登录成功")

}
