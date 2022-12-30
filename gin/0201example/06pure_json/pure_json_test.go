package pure_json

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 通常, json使用unicode替换特殊字符, 例如<变为\u003c
// 如果要按字面对这些字符进行编码, 可以使用PureJSON
func Test_Pure_Json(t *testing.T) {
	router := gin.Default()

	// 提供unicode实体 ==> {"html":"\u003cbr\u003ehello, world!\u003c/br\u003e"}
	router.GET("json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"html": "<br>hello, world!</br>",
		})
	})

	// 提供字面字符 ==> {"html":"<br>hello, world!</br>"}
	router.GET("/pure-json", func(ctx *gin.Context) {
		ctx.PureJSON(http.StatusOK, gin.H{
			"html": "<br>hello, world!</br>",
		})
	})

	router.Run(":8080")
}
