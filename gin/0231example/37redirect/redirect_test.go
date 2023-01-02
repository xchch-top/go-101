package _7redirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 重定向
func Test_Redirect(t *testing.T) {
	router := gin.Default()

	router.GET("/search", func(ctx *gin.Context) {
		// 301 永久性转移
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	router.POST("/search2", func(ctx *gin.Context) {
		// 302 暂时性转移
		ctx.Redirect(http.StatusFound, "https://www.baidu.com/")
	})

	router.Run(":8080")
}

// 路由重定向, 使用HandleContext
func Test_Redirect_V2(t *testing.T) {
	router := gin.Default()

	// 转发, 响应的http状态码是200
	router.GET("/test", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/test2"
		router.HandleContext(ctx)
	})

	router.GET("/test2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"hello": "world"})
	})

	router.Run(":8080")
}
