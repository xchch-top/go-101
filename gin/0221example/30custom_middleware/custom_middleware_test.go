package custom_middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
	"time"
)

// 自定义中间件
func Test_Custom_Middleware(t *testing.T) {
	router := gin.New()
	router.Use(Logger())

	router.GET("/mid", func(ctx *gin.Context) {
		time.Sleep(1 * time.Second)
		example := ctx.MustGet("example").(string)
		ctx.String(http.StatusOK, example)
	})

	router.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		// 请求前
		// 设置example变量
		ctx.Set("example", "12345")

		ctx.Next()

		// 请求后
		latency := time.Since(t)
		status := ctx.Writer.Status()
		log.Printf("response status: %v, 耗时: %vms", status, latency.Milliseconds())
	}
}
