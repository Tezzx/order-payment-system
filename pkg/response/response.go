package response

import "github.com/gin-gonic/gin"

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": data,
	})
}

// Error 失败响应
func Error(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"msg":  msg,
		"data": nil,
	})
}
