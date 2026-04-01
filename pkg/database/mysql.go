package database

import (
	"fmt"
	"order-payment-system/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// *gorm.DB使用时表现为一个句柄，实际内部维护着一个连接池
func InitMySQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
