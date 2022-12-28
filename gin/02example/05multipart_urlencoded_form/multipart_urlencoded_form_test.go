package multipart_urlencoded_form

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// Multipart Urlencoded表单
func Test_Multipart_Urlencoded_Form(t *testing.T) {
	router := gin.Default()

	router.POST("/form_post", func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		nick := ctx.DefaultPostForm("nick", "anonymous")

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.Run(":8080")
}
