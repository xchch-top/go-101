package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_Http_Method(t *testing.T) {
	// 禁用控制台颜色, main方法中能体现效果
	// gin.DisableConsoleColor()
	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	router.GET("/someGet", httpFunc)
	router.POST("/somePost", httpFunc)
	router.PUT("/somePut", httpFunc)
	router.DELETE("/someDelete", httpFunc)
	router.PATCH("/somePatch", httpFunc)
	router.HEAD("/someHead", httpFunc)
	router.OPTIONS("/someOptions", httpFunc)

	router.Run(":8080")
}

func httpFunc(ctx *gin.Context) {
	ctx.String(http.StatusOK, ctx.Request.RequestURI)
}
