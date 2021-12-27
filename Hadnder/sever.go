package Hander

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.Default()
	r.Use(cors.Default())

	//配置静态文件
	r.Static("/", "Final_Project/Static")

	//设置路由组
	linx := r.Group("linx")
	{
		linx.POST("/transfer", Transfer)
		linx.POST("/cAmount", GetContractAmount)
		linx.POST("/getLimit", Limit)
		linx.POST("/sendAll", SendAll)
		linx.POST("/changer", Chager)
		linx.POST("/freedLimit", FreeLimit)
	}
	err := r.Run("127.0.0.1:1234")
	return err
}

//封装函数，用于向客户端返回正确消息
func respOK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}
