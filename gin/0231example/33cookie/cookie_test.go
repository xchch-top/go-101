package cookie

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"testing"
)

// 设置和获取 Cookie
func Test_Cookie(t *testing.T) {
	router := gin.Default()

	router.GET("/cookie", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("gin_cookie")
		if err != nil {
			cookie = "not set"
			ctx.SetCookie("gin_cookie", uuid.New().String(), 3600, "/", "localhost", false, true)
		}
		log.Printf("cookie value: %s \n", cookie)
	})

	router.Run(":8080")
}
