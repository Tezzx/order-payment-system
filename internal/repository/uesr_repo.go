package repository

import (
	"order-payment-system/internal/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(user *model.User) error {
	err := u.db.Create(user).Error
	return err
}

func (u *UserRepo) GetByUsername(username string) (string, error) {
	var user model.User
	err := u.db.Where("username=?", username).Select("password").First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Password, nil
}
