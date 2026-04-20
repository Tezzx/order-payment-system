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

// 创建用户
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
func (u *UserRepo) CheckUsernameExists(username string) (bool, error) {
	var count int64
	err := u.db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (u *UserRepo) GetID(username string) (uint, error) {
	var userID uint
	err := u.db.Model(&model.User{}).
		Select("id").
		Where("username = ?", username).
		Scan(&userID).
		Error
	if err != nil {
		return 0, err
	}
	return userID, nil
}
