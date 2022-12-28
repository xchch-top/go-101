package html_render

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"testing"
	"time"
)

// 使用LoadHTMLGlob()或者LoadHTMLFiles()渲染html页面
func Test_Html_Render(t *testing.T) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**")

	// r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}

// 使用不同目录下名称相同的模板
func Test_Html_Render_V2(t *testing.T) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	router.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	router.Run(":8080")
}

// 自定义模板渲染器
func Test_Html_Render_V3(t *testing.T) {
	router := gin.Default()
	html := template.Must(template.ParseFiles("file1", "file2"))
	router.SetHTMLTemplate(html)

	router.Run(":8080")
}

// 自定义模板功能
func Test_Html_Render_V5(t *testing.T) {
	router := gin.Default()
	// 自定义分隔符
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./testdata/template/raw.tmpl")

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
