package log_format

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
	"time"
)

// 自定义日志文件
func Test_Log_Format(t *testing.T) {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP, params.TimeStamp.Format(time.RFC1123), params.Method, params.Path,
			params.Request.Proto, params.StatusCode, params.Latency, params.Request.UserAgent(),
			params.ErrorMessage)
	}))

	router.Use(gin.Recovery())

	router.GET("ping", func(ctx *gin.Context) {
		log.Println("ping method call")
		ctx.String(http.StatusOK, "pong")
	})

	router.Run(":8080")
}
