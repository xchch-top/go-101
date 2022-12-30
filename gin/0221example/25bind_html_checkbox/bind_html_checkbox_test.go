package bind_html_checkbox

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 绑定html复选框
func Test_Bind_Html_Checkbox(t *testing.T) {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "form.html", nil)
	})

	router.POST("/", func(ctx *gin.Context) {
		var fakeForm myForm
		ctx.Bind(&fakeForm)
		ctx.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
	})

	router.Run(":8080")
}

type myForm struct {
	Colors []string `form:"colors[]"`
}
