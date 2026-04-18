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

func NewUserService(userrepo *repository.UserRepo) *UserService {
	return &UserService{
		userRepo: userrepo,
	}
}

// 注册
func (u *UserService) RegisterUser(userName, userPassword string) error {
	bol, err := u.userRepo.CheckUsernameExists(userName)
	if err != nil || bol == true {
		return errors.New("用户已存在")
	}
	userPassword, err = util.HashPassword(userPassword)
	if err != nil {
		return err
	}
	user := model.User{
		Username: userName,
		Password: userPassword,
	}
	return u.userRepo.CreateUser(&user)
}

// 登录
func (u *UserService) LoginUser(userName, userPassword string) error {
	password, err := u.userRepo.GetByUsername(userName)
	if err != nil {
		return err
	}
	return util.VerifyPassword(userPassword, password)
}

func (u *UserService) TokenCreate(userName string) (string, error) {
	token, err := jwt.GenerateJWT(userName)
	return token, err
}
