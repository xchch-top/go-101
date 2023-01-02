package multipart_urlencoded_binding

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// Multipart Urlencoded绑定
func Test_Multipart_Urlencoded_Binding(t *testing.T) {
	router := gin.Default()

	router.POST("/login", func(ctx *gin.Context) {
		// 你可以使用显式声明绑定 multipart form;
		// ctx.ShouldBindWith(&form, binding.Form)
		// 或者简单使用ShouldBind方法自动绑定, 将自动选择合适的绑定
		var form LoginForm
		if ctx.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				ctx.JSON(http.StatusOK, gin.H{"status": "you are login in"})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"status": "you are login in"})
			}
		}
	})

	router.Run(":8080")
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
