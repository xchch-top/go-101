package query

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

// 同时使用Query查询和Form查询
func Test_Query(t *testing.T) {
	router := gin.Default()

	router.POST("/post", func(ctx *gin.Context) {
		id := ctx.Query("id")
		page := ctx.DefaultQuery("page", "0")
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")

		// id: 1234, page: 1, name: manu, message: this_is_great
		fmt.Printf("id: %v, page: %v, name: %v, message: %v\n", id, page, name, message)
	})

	router.Run(":8080")
}

// POST /post?id=1234&page=1 HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
//
// name=manu&message=this_is_great
