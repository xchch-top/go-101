package querystring_postform

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

// 映射查询字符串或表单参数
func Test_QueryString_Postform(t *testing.T) {
	router := gin.Default()

	// http request:
	// curl --location -g --request POST 'http://localhost:8080/post?ids[a]=a1&ids[b]=b2' \
	// --header 'Content-Type: application/x-www-form-urlencoded' \
	// --data-urlencode 'names[first]=f' --data-urlencode 'names[second]=s'
	router.POST("/post", func(ctx *gin.Context) {
		ids := ctx.QueryMap("ids")
		names := ctx.PostFormMap("names")
		// ids: map[a:a1 b:b2], names: map[first:f second:s]
		fmt.Printf("ids: %v, names: %v\n", ids, names)
	})

	router.Run(":8080")
}
