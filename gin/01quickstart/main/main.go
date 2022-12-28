package main

import "github.com/gin-gonic/gin"

func main() {
	// default方法中使用了Logger和Recovery中间件
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		// H是map[string]interface{}的别名
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run()
}
