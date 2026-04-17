package response

import "github.com/gin-gonic/gin"

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

// Error 失败响应
func Error(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  msg,
		"data": nil,
	})
}
