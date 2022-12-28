package http_server_push

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"testing"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

// http2 server推送
func Test_Http2_Server_Push(t *testing.T) {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(html)

	router.GET("/", func(ctx *gin.Context) {
		pusher := ctx.Writer.Pusher()
		if pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			err := pusher.Push("/assets/app.js", nil)
			if err != nil {
				log.Printf("failed to push: %v", err)
			}

			ctx.HTML(http.StatusOK, "https", gin.H{
				"status": "success",
			})
		}
	})

	router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
