package querystring

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 查询字符串参数
func Test_QueryString(t *testing.T) {
	router := gin.Default()

	router.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "f")
		lastname := ctx.Query("lastname")

		ctx.String(http.StatusOK, fmt.Sprintf("hello %s %s", firstname, lastname))
	})

	router.Run(":8080")
}
