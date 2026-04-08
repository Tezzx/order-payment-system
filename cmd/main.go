package main

import (
	"fmt"
	"order-payment-system/config"
	"order-payment-system/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//获取数据库连接
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}
	db, err := database.InitMySQL(&cfg.Database)
	if err != nil {
		return
	}
	_ = db

	//路由引擎
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		panic(fmt.Sprintf("启动服务失败: %v", err))
	}
}
