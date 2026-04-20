package main

import (
	"fmt"
	"order-payment-system/config"
	"order-payment-system/internal/handler"
	"order-payment-system/internal/model"
	"order-payment-system/internal/repository"
	"order-payment-system/internal/service"
	"order-payment-system/pkg/database"
	"order-payment-system/pkg/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	//读取配置文件
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("配置文件读取失败")
		return
	}
	port := strconv.Itoa(cfg.Server.Port)

	//连接数据库
	db, err := database.InitMySQL(&cfg.Database)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	fmt.Println("数据库连接成功")

	//自动建表
	err = db.AutoMigrate(&model.User{}, &model.Goods{})
	if err != nil {
		fmt.Println("数据表创建失败:", err)
		return
	}

	//依赖注入
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	goodsRepo := repository.NewGoodsRepo(db)
	goodsService := service.NewGoodsService(goodsRepo)
	goodsHandler := handler.NewGoodsHandler(goodsService)
	//商品数据库初始化
	goodsHandler.GoodsInitial()

	r := gin.Default()

	r.Use(middleware.CorsMiddleware())
	//路径是根据工作路径来算而不是main.go所在路径
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"goodsName": "异色拉特",
			"price":     99.00,
		})
	})
	r.GET("/auth", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/register", userHandler.RegisterUser)
	r.POST("/login", userHandler.LoginUser)

	order := r.Group("/order")
	{
		order.POST("/create", middleware.TokenIdentify(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "下单成功！等待支付",
			})
			//创建订单
		})

		order.GET("/pay", func(c *gin.Context) {
			c.HTML(200, "pay.html", gin.H{
				"price": 99.00,
			})
		})
	}

	r.Run(":" + port)
}
