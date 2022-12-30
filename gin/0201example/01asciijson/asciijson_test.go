package asciijson

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 使用AsciiJSON生成具有转义的非ASCII字符的ASCII-only JSON
func Test_Asciijson(t *testing.T) {
	router := gin.Default()

	router.GET("/someJson", func(ctx *gin.Context) {
		data := map[string]any{
			"lang": "Go语言",
			"tag":  "<br>",
		}

		// 输出 {"lang":"Go\u8bed\u8a00","tag":"\u003cbr\u003e"}
		ctx.AsciiJSON(http.StatusOK, data)
	})

	// 监听在0.0.0.0:8080上启动服务
	router.Run(":8080")
}
