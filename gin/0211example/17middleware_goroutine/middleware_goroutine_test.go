package middleware_goroutine

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
	"time"
)

// 在中间件或handler中启动新的goroutine时, 不能使用原始的上下文, 必须使用只读副本
func Test_Middleware_Goroutine(t *testing.T) {
	router := gin.Default()

	router.GET("/long-async-1", func(ctx *gin.Context) {
		// 创建副本
		cCopy := ctx.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// 使用副本
			log.Println("Done, in path " + cCopy.Request.URL.Path)
			// cCopy.String(http.StatusOK, "cCopy") 方法panic
		}()
		ctx.String(http.StatusOK, "ctx")
	})

	router.GET("/long-async-2", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done, in path " + ctx.Request.URL.Path)
		ctx.String(http.StatusOK, "ctx")
	})

	router.Run(":8080")
}
