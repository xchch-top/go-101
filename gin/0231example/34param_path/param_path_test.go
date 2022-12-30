package param_path

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// http://localhost:8080/user/john ==> 命中handler1
// http://localhost:8080/user/john/ ==> 命中handler2
// http://localhost:8080/user/john/send ==> 命中handler2
func Test_Param_Path(t *testing.T) {
	router := gin.Default()

	// 此handler将匹配/user/john, 但不会匹配/user或者/user/
	router.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello %s", name)
	})

	// 此handler将匹配/user/john/和/user/john/send
	router.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + "is" + action
		ctx.String(http.StatusOK, message)
	})

	router.Run(":8080")
}

// http://localhost:8080/user/john ==> 命中
// http://localhost:8080/user/jhon/ ==> 301 http://localhost:8080/user/jhon
// http://localhost:8080/user/jhon/send ==> 404
func Test_Param_Path_V2(t *testing.T) {
	router := gin.Default()

	// 此handler将匹配/user/john, 但不会匹配/user或者/user/
	router.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello %s", name)
	})

	router.Run(":8080")
}

// http://localhost:8080/user/john ==> 301 http://localhost:8080/user/jhon/
// http://localhost:8080/user/jhon/ ==> 命中
// http://localhost:8080/user/jhon/send ==> 命中
func Test_Param_Path_V3(t *testing.T) {
	router := gin.Default()

	// 此handler将匹配/user/john/和/user/john/send
	// 如果没有其他路由匹配到/user/john, 它将重定向到/user/john/
	router.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + "is" + action
		ctx.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
