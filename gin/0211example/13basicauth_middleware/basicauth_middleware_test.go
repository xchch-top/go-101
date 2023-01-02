package basicauth_middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"}}

// 使用BasicAuth中间件
func Test_BasicAuth_Middleware(t *testing.T) {
	router := gin.Default()

	// 路由组使用gin.BasicAuth()中间件
	// gin.Accounts是map[string]string的别名
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets端点 ==> 触发 http://localhost:8080/admin/secrets
	authorized.GET("/secrets", func(ctx *gin.Context) {
		// 获取用户, 它是由BasicAuth中间件设置的
		user := ctx.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "Not Found"})
		}
	})

	router.Run(":8080")
}

// curl --location --request GET 'http://localhost:8080/admin/secrets' --header 'Authorization: Basic Zm9vOmJhcg=='
