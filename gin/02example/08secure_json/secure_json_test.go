package secure_json

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 使用SecureJson防止json劫持, 如果给定的结构是数组值, 则默认预置 while(1) 到响应体
func Test_Secure_JSON(t *testing.T) {
	router := gin.Default()
	// 你也可以声依永自己的SecureJson前缀
	// router.SecureJsonPrefix(")]},\n" )

	router.GET("/secure-json", func(ctx *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// 输出 while(1);["lena","austin","foo"]
		ctx.SecureJSON(http.StatusOK, names)
	})

	router.GET("/secure-json2", func(ctx *gin.Context) {
		h := gin.H{
			"name": "json2",
		}
		// 输出 {"name":"json2"}
		ctx.SecureJSON(http.StatusOK, h)
	})

	router.Run(":8080")
}
