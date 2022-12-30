package define_log_format

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

// 定义路由日志的格式
func Test_Define_Long_format(t *testing.T) {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		// 使用标准包
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	router.Run(":8080")
}

// before: [GIN-debug] GET    /ping                     --> gitlab.xchch.top/golang-group/go-101/gin/02example/19define_log_format.Test_Define_Long_format.func1 (3 handlers)
// after:  2022/12/29 19:47:40 endpoint GET /ping gitlab.xchch.top/golang-group/go-101/gin/02example/19define_log_format.Test_Define_Long_format.func2 3
