package service

import (
	"errors"
	"order-payment-system/internal/model"
	"order-payment-system/internal/repository"
	"order-payment-system/pkg/jwt"
	"order-payment-system/pkg/util"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// 注册
func (u *UserService) RegisterUser(userName, userPassword string) (uint, error) {
	bol, err := u.userRepo.CheckUsernameExists(userName)
	if err != nil || bol == true {
		return 0, errors.New("用户已存在")
	}
	userPassword, err = util.HashPassword(userPassword)
	if err != nil {
		return 0, err
	}
	user := model.User{
		Username: userName,
		Password: userPassword,
	}
	userID, err := u.userRepo.GetID(userName)
	if err != nil {
		return 0, err
	}
	return userID, u.userRepo.CreateUser(&user)
}

// 登录
func (u *UserService) LoginUser(userName, userPassword string) (uint, error) {
	password, err := u.userRepo.GetByUsername(userName)
	if err != nil {
		return 0, err
	}
	userID, err := u.userRepo.GetID(userName)
	if err != nil {
		return 0, err
	}
	return userID, util.VerifyPassword(userPassword, password)
}

func (u *UserService) TokenCreate(userID uint) (string, error) {
	token, err := jwt.GenerateJWT(userID)
	return token, err
}
