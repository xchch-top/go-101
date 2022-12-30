package write_log

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"testing"
)

// 如何记录日志
func Test_Write_Log(t *testing.T) {
	// 禁用控制台颜色, 将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()

	// 记录到文件
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台, 请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	router.Run(":8080")
}
