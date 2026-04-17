package service

import (
	"order-payment-system/internal/model"
	"order-payment-system/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(userrepo *repository.UserRepo) *UserService {
	return &UserService{
		userRepo: userrepo,
	}
}

// 注册
func (u *UserService) RegisterUser(userName, userPassword string) error {
	user := model.User{
		Username: userName,
		Password: userPassword,
	}
	return u.userRepo.CreateUser(&user)
}

// 登录
func (u *UserService) LoginUser(userName, userPassword string) string {
	password, err := u.userRepo.GetByUsername(userName)
	var s string
	if err != nil {
		s = "用户不存在"
		return s
	}
	if password != userPassword {
		s = "账号或密码错误"
		return s
	}
	return ""
}
