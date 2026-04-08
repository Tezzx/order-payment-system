package repository

import (
	"order-payment-system/internal/model"

	"gorm.io/gorm"
)

// 数据库增删改查
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(d *gorm.DB) *UserRepository {
	return &UserRepository{db: d}
}

func (this *UserRepository) AddUser(user *model.User) error {
	return this.db.Create(user).Error
}

func (this *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := this.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
