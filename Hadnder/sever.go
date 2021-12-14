package Hander

import (
	"github.com/gin-gonic/gin"
)

func start() error{
	r := gin.Default()
	//配置静态文件
	r.Static("/", "Final_Project/Static")

	//设置路由组
	linx := r.Group("linx")
	{
		linx.POST("/transfer",Transfer)
		linx.POST("/CAmont",GetContractAmount)
		linx.POST("/GetLimit",Limit)
	}
	err := r.Run()
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
